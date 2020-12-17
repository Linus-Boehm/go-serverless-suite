package mailings

type Service struct {
	provider IMailProvider
}

type Renderer interface {
	Render(data interface{}) (*string, error)
}

func NewMailingsService(provider IMailProvider) *Service {
	return &Service{provider: provider}
}

func (m *Service) SendSimpleContactForm(input SimpleContactFormInput, renderer Renderer) error {
	content, err := renderer.Render(input)
	if err != nil {
		return err
	}
	mmi := MinimumMailInput{
		FromMail:    input.ToMail,
		ToMail:      input.ToMail,
		Subject:     &input.Subject,
		HTMLContent: *content,
	}

	return m.provider.SendSingleMail(mmi)
}
