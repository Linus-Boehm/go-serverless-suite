package mailings

import "github.com/Linus-Boehm/go-serverless-suite/templates"

type MailSender interface {
	SendSimpleContactForm(input SimpleContactFormInput, renderer Renderer) error
}

type service struct {
	provider MailProvider
}

type Renderer interface {
	Render(data interface{}) (*string, error)
	RenderWithHTML(data interface{}) (*templates.HTMLTemplate, error)
}

func NewMailSender(provider MailProvider) MailSender {
	return &service{provider: provider}
}

func (m *service) SendSimpleContactForm(input SimpleContactFormInput, renderer Renderer) error {
	content, err := renderer.RenderWithHTML(input)
	if err != nil {
		return err
	}
	mmi := MinimumMailInput{
		FromMail:     input.SenderMail,
		ToMail:       input.ToMail,
		Subject:      &input.Subject,
		HTMLTemplate: *content,
	}

	return m.provider.SendSingleMail(mmi)
}
