package itf

import (
	"github.com/Linus-Boehm/go-serverless-suite/entity"
)

//go:generate mockgen -destination=mailings_mocks.go -package=itf -source=mailings.go

type TplRenderer interface {
	GetRaw() *string
	Render(data interface{}) (*string, error)
	RenderWithHTML(data interface{}) (*entity.HTMLTemplate, error)
}

type Mailer interface {
	SendSimpleContactForm(input entity.ContactForm, renderer TplRenderer) error
	GetContactLists() ([]entity.MailContactList, error)
	GetProvider() MailerProvider
}

type MailerProvider interface {
	SendSingleMail(input entity.MinimalMail) error
	GetContactLists() ([]entity.MailContactList, error)
}
