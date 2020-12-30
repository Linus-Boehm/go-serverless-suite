package mailer

import (
	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
)

type service struct {
	provider itf.MailerProvider
}

func NewMailerService(provider itf.MailerProvider) itf.Mailer {
	return &service{provider: provider}
}

func (m *service) SendSimpleContactForm(input entity.ContactForm, renderer itf.TplRenderer) error {
	content, err := renderer.RenderWithHTML(input)
	if err != nil {
		return err
	}
	mmi := entity.MinimalMail{
		FromMail:     input.SenderMail,
		ToMail:       input.ToMail,
		Subject:      &input.Subject,
		HTMLTemplate: *content,
	}

	return m.provider.SendSingleMail(mmi)
}

func (m *service) GetContactLists() ([]entity.MailContactList, error) {
	return m.provider.GetContactLists()
}

func (m *service) GetProvider() itf.MailerProvider {
	return m.provider
}
