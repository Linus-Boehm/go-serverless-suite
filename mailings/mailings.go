package mailings

type MailSender interface {
	SendSimpleContactForm(input SimpleContactFormInput, renderer Renderer) error
}

type service struct {
	provider MailProvider
}

type Renderer interface {
	Render(data interface{}) (*string, error)
}

func NewMailingsService(provider MailProvider) MailSender {
	return &service{provider: provider}
}

func (m *service) SendSimpleContactForm(input SimpleContactFormInput, renderer Renderer) error {
	content, err := renderer.Render(input)
	if err != nil {
		return err
	}
	mmi := MinimumMailInput{
		FromMail:    input.SenderMail,
		ToMail:      input.ToMail,
		Subject:     &input.Subject,
		HTMLContent: *content,
	}

	return m.provider.SendSingleMail(mmi)
}
