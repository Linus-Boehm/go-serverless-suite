package itf

import "github.com/Linus-Boehm/go-serverless-suite/entity"

//go:generate mockgen -destination=crm_mocks.go -package=itf -source=crm.go

type NewsWriter interface {
}

type CRMServicer interface {
	GetMailer() Mailer
	CreateNewUser(user entity.User) (entity.User, error)
	CreateSubscription(subs entity.CRMEmailListSubscription) (entity.CRMEmailListSubscription, error)
	CreateSubscriptions(subscriptions []entity.CRMEmailListSubscription) ([]entity.CRMEmailListSubscription, error)
	SendDoubleOptInMail(options entity.CRMOptInMailOptions, renderer TplRenderer) error
	ValidateEmail(email entity.ID, userID entity.ID, subID entity.ID) error
}

type CRMProvider interface {
	GetSubscriptionsOfEmail(email entity.ID) ([]entity.CRMEmailListSubscription, error)
	PutSubscription(entity.CRMEmailListSubscription) error
	PutSubscriptions([]entity.CRMEmailListSubscription) error
	GetSubscriptionsOfList(listID entity.ID) ([]entity.CRMEmailListSubscription, error)
}
