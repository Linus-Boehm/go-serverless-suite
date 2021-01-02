package itf

import "github.com/Linus-Boehm/go-serverless-suite/entity"

type NewsWriter interface {
}

type CRMServicer interface {
	GetMailer() Mailer
	CreateSubscription(subs []entity.CRMEmailListSubscription, confirmationTPL entity.HTMLTemplate) error
	ValidateEmail(email entity.ID) error

}

type CRMProvider interface {
	GetSubscriptionsOfEmail(email entity.ID) ([]entity.CRMEmailListSubscription, error)
	PutSubscription(entity.CRMEmailListSubscription) error
	PutSubscriptions([]entity.CRMEmailListSubscription) error
	GetSubscriptionsOfList(listID entity.ID) ([]entity.CRMEmailListSubscription, error)
}
