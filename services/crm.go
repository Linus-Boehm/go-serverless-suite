package services

import (
	"fmt"
	"html/template"
	"net/url"

	"github.com/Linus-Boehm/go-serverless-suite/common"
	"github.com/Linus-Boehm/go-serverless-suite/common/tplengine"
	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
)

type CRMService struct {
	mailer     itf.Mailer
	repo       itf.CRMProvider
	senderMail entity.Mail
	userRepo   itf.UserProvider
}

func NewCRMService(mailer itf.Mailer, repo itf.CRMProvider, userRepo itf.UserProvider, senderMail entity.Mail) *CRMService {
	return &CRMService{
		mailer:     mailer,
		repo:       repo,
		senderMail: senderMail,
		userRepo:   userRepo,
	}
}

func (c CRMService) GetMailer() itf.Mailer {
	return c.mailer
}
func (c CRMService) CreateNewUser(user entity.User) (entity.User, error) {
	user.ID.NewV4IfEmpty()
	user.Timestamps.CreatedNow()
	err := c.userRepo.PutUser(user)
	return user, err
}

// CreateSubscription creates a new subscription, if the email is already subscribed to the ListID, then no new subscription is created, but the current one is returned
func (c CRMService) CreateSubscription(subscription entity.CRMEmailListSubscription) (entity.CRMEmailListSubscription, error) {
	currentSubs, err := c.repo.GetSubscriptionsOfEmail(subscription.EMail)
	if err != nil {
		return subscription, err
	}
	// Check if user is already subscribed, if yes return subscription
	for _, userSub := range currentSubs {
		if userSub.ListID == subscription.ListID {
			return userSub, nil
		}
	}

	subscription.Timestamps.CreatedNow()
	subscription.SubscriptionID.NewV4IfEmpty()

	return subscription, c.repo.PutSubscription(subscription)
}

func (c CRMService) CreateSubscriptions(subscriptions []entity.CRMEmailListSubscription) ([]entity.CRMEmailListSubscription, error) {
	if len(subscriptions) == 0 {
		return nil, nil
	}
	currentSubs, err := c.repo.GetSubscriptionsOfEmail(subscriptions[0].EMail)
	if err != nil {
		return nil, err
	}
	subsToCreate := []entity.CRMEmailListSubscription{}
	// get only subscriptions which are not already
	for _, sub := range subscriptions {
		found := false
		for _, currentSub := range currentSubs {
			// if user already exists, only perform update when not already opted in
			if sub.ListID == currentSub.ListID {
				if currentSub.Status.String() != entity.UserOptedInSubscriptionStatus.String() {
					sub.Status = currentSub.Status
					sub.Timestamps = currentSub.Timestamps
					sub.Timestamps.UpdatedNow()
				} else {
					found = true
				}
			}
		}
		if !found {
			subsToCreate = append(subsToCreate, sub)
		}
	}

	return subsToCreate, c.repo.PutSubscriptions(subsToCreate)
}

func (c CRMService) SendDoubleOptInMail(options entity.CRMOptInMailOptions, renderer itf.TplRenderer) error {
	if renderer == nil {
		engine, err := tplengine.LoadLayoutTemplate(tplengine.CRMOptInMailManifest)
		if err != nil {
			return err
		}
		renderer = engine
	}

	encodedEmail := url.QueryEscape(options.EMail.String())
	confirmURL := fmt.Sprintf("%s?id=%s&subid=%s&email=%s", options.ConfirmationPath, options.UserID, options.SubID, encodedEmail)
	htmlLink := template.HTML(fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, confirmURL, confirmURL))
	tplOptions := entity.CRMOptInMailTemplateOptions{
		ConfirmURL: htmlLink,
		Email:      options.EMail.String(),
		FullName:   options.Fullname,
	}
	htmlTemplate, err := renderer.RenderWithHTML(tplOptions)
	if err != nil {
		return err
	}

	var subject *string
	if options.Subject != nil {
		subject = options.Subject
	} else {
		subject = common.StringPtr("Please confirm your Sign-up")
	}
	mailConfig := entity.MinimalMail{
		FromMail: c.senderMail,
		ToMail: entity.Mail{
			Name: options.Fullname,
			Mail: options.EMail.String(),
		},
		Subject:      subject,
		HTMLTemplate: *htmlTemplate,
	}
	return c.mailer.GetProvider().SendSingleMail(mailConfig)
}

func (c CRMService) ValidateEmail(email entity.ID, userID entity.ID, subID entity.ID) error {
	subs, err := c.repo.GetSubscriptionsOfEmail(email)
	if err != nil {
		return err
	}
	user, err := c.userRepo.ReadUser(userID)
	if err != nil {
		return err
	}
	var toConfirm []entity.CRMEmailListSubscription
	listIDs := []entity.ID{}
	//Filter for subscription
	for _, sub := range subs {
		if sub.SubscriptionID.String() == subID.String() {
			// Set user to confirmed in our DB
			sub.Status = entity.UserOptedInSubscriptionStatus
			sub.UpdatedNow()
			toConfirm = append(toConfirm, sub)

		}
		//TODO move within if statement to only update list that are not already subscribed
		listIDs = append(listIDs, sub.ListID)
	}

	mailProvider := c.mailer.GetProvider()
	//TODO if user is already verified only update user
	err = mailProvider.CreateUser(user, listIDs)
	if err != nil {
		return err
	}

	if err := c.repo.PutSubscriptions(toConfirm); err != nil {
		return err
	}
	if !user.EmailVerified {
		user.EmailVerified = true
		user.UpdatedNow()
		return c.userRepo.PutUser(user)
	}
	return nil

}
