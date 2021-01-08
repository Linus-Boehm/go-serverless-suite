package services

import (
	"fmt"
	"github.com/Linus-Boehm/go-serverless-suite/common"
	"github.com/Linus-Boehm/go-serverless-suite/common/tplreader"
	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
	"html/template"
	"net/url"
)

type crmSVC struct {
	mailer itf.Mailer
	repo itf.CRMProvider
	senderMail entity.Mail
	userRepo itf.UserProvider
}

func NewCRMService(mailer itf.Mailer, repo itf.CRMProvider, senderMail entity.Mail) *crmSVC {
	return &crmSVC{
		mailer: mailer,
		repo: repo,
		senderMail: senderMail,
	}
}

func (c crmSVC) GetMailer() itf.Mailer {
	panic("implement me")
}
func (c crmSVC) CreateNewUser(user entity.User) (entity.User,error) {
	user.ID.NewV4IfEmpty()
	user.Timestamps.CreatedNow()
	err := c.userRepo.PutUser(user)
	return user, err
}

// CreateSubscription creates a new subscription, if the email is already subscribed to the ListID, then no new subscription is created, but the current one is returned
func (c crmSVC) CreateSubscription(subscription entity.CRMEmailListSubscription) (entity.CRMEmailListSubscription, error) {
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

func (c crmSVC) SendDoubleOptInMail(options entity.CRMOptInMailOptions ) error {
	if options.FS == nil {
		options.FS = &tplreader.DefaultManifests
	}
	if options.Template == nil {
		options.Template = &tplreader.CRMOptInMailManifest
	}
	reader, err :=  tplreader.LoadCustomTemplate(options.FS, *options.Template)
	if err != nil {
		return err
	}
	encodedEmail := url.QueryEscape(options.EMail.String())
	confirmUrl := fmt.Sprintf("%s/?id=%s&subid=%s&email=%s", options.ConfirmationPath, options.UserID, options.SubID, encodedEmail)
	htmlLink := template.HTML(fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, confirmUrl, confirmUrl))
	tplOptions := entity.CRMOptInMailTemplateOptions{
		ConfirmURL:   htmlLink,
		Email:        options.EMail.String(),
		FullName:     options.Fullname,
	}
	htmlTemplate, err := reader.RenderWithHTML(tplOptions)
	if err != nil {
		return err
	}

	var subject *string
	if options.Subject != nil {
		subject = options.Subject
	}else {
		subject = common.StringPtr("Please confirm your Sign-up")
	}
	mailConfig := entity.MinimalMail{
		FromMail:     c.senderMail,
		ToMail:       entity.Mail{
			Name: options.Fullname,
			Mail: options.EMail.String(),
		},
		Subject:      subject,
		HTMLTemplate: *htmlTemplate,
	}
	return c.mailer.GetProvider().SendSingleMail(mailConfig)
}

func (c crmSVC) ValidateEmail(email entity.ID, userID entity.ID, subID entity.ID) error {
	subs, err := c.repo.GetSubscriptionsOfEmail(email)
	if err != nil {
		return err
	}
	user, err := c.userRepo.ReadUser(userID)
	if err != nil {
		return err
	}
	var toConfirm []entity.CRMEmailListSubscription
	//Filter for subscription
	for _, sub := range subs {
		if sub.SubscriptionID.String() == subID.String(){
			// Set user to confirmed in our DB
			sub.Status = entity.UserOptedInSubscriptionStatus
			sub.UpdatedNow()
			toConfirm = append(toConfirm, sub)
		}
	}
	//TODO add user here to provider email list

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


