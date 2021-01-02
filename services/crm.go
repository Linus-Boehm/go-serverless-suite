package services

import (
	"github.com/Linus-Boehm/go-serverless-suite/common"
	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
	"github.com/pkg/errors"
)

type crmSVC struct {
	mailer itf.Mailer
	repo itf.CRMProvider
	senderMail entity.Mail
}

func NewCRMService(mailer itf.Mailer, repo itf.CRMProvider, senderMail entity.Mail) itf.CRMServicer {
	return &crmSVC{
		mailer: mailer,
		repo: repo,
		senderMail: senderMail,
	}
}

func (c crmSVC) GetMailer() itf.Mailer {
	panic("implement me")
}

func (c crmSVC) CreateSubscription(subscriptions []entity.CRMEmailListSubscription, confirmationTPL entity.HTMLTemplate) error {
	if len(subscriptions) == 0 {
		return nil
	}
	subID := entity.NewEntityIDV4()
	email := subscriptions[0].EMail
	for i,sub := range subscriptions {
		if sub.EMail.String() != email.String() {
			return errors.New("provided different emails for subscription batch")
		}
		subscriptions[i].Timestamps.CreatedNow()
		subscriptions[i].SubscriptionID = subID
	}
	if err := c.repo.PutSubscriptions(subscriptions); err != nil {
		return err
	}
	mail := entity.MinimalMail{
		FromMail:     c.senderMail,
		ToMail:       entity.Mail{
			Mail: email.String(),
		},
		Subject:      common.StringPtr(confirmationTPL.Title),
		HTMLTemplate: confirmationTPL,
	}
	return c.mailer.GetProvider().SendSingleMail(mail)
}

func (c crmSVC) ValidateEmail(email entity.ID) error {
	panic("implement me")
}


