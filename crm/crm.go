package crm

import "github.com/Linus-Boehm/go-serverless-suite/templates"

type CRMServicer interface {
	SendSimpleContactForm(input SimpleContactFormInput, renderer Renderer) error
	GetContactLists() ([]ContactList, error)
}

type service struct {
	provider CRMProvider
}

type Renderer interface {
	Render(data interface{}) (*string, error)
	RenderWithHTML(data interface{}) (*templates.HTMLTemplate, error)
}

func NewCRMService(provider CRMProvider) CRMServicer {
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

func (m *service) GetContactLists() ([]ContactList, error) {
	return m.provider.GetContactLists()
}
