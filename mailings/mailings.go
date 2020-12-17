package mailings

type MailingsService struct {
	provider IMailProvider
}

type Renderer interface {
	Render(data interface{}) (*string, error)
}

func NewMailingsService(provider IMailProvider) *MailingsService {
	return &MailingsService{provider: provider}
}


func (m *MailingsService) SendSimpleContactForm(input SimpleContactFormInput, renderer Renderer) error {
	content, err := renderer.Render(input)
	if err != nil {
		return err
	}
	mmi := MinimumMailInput{
		FromMail:    input.ToMail,
		ToMail:      input.ToMail,
		Subject:     &input.Subject,
		HtmlContent: *content,
	}

	return m.provider.SendSingleMail(mmi)
}