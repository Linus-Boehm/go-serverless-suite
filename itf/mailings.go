package itf

import (
	"io/fs"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
)

//go:generate mockgen -destination=mailings_mocks.go -package=itf -source=mailings.go

type TplRenderer interface {
	Render(data interface{}) (*string, error)
	RenderWithHTML(data interface{}) (*entity.HTMLTemplate, error)
	WithTemplate(fs fs.FS, manifest entity.TemplateManifest, tplName string) (TplRenderer, error)
}

type Mailer interface {
	SendSimpleContactForm(input entity.ContactForm, renderer TplRenderer) error
	GetContactLists() ([]entity.MailContactList, error)
	GetProvider() MailerProvider
}

type MailerProvider interface {
	SendSingleMail(input entity.MinimalMail) error
	GetContactLists() ([]entity.MailContactList, error)
	GetDefaultSender() *entity.Mail
	CreateUser(user entity.User, listIDs []entity.ID) error
}
