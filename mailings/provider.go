package mailings

import (
	"github.com/Linus-Boehm/go-serverless-suite/templates"
)

type MailProvider interface {
	SendSingleMail(input MinimumMailInput) error
}

type Mail struct {
	Name string
	Mail string
}

type MinimumMailInput struct {
	FromMail     Mail
	ToMail       Mail
	Subject      *string
	HTMLTemplate templates.HTMLTemplate
}

func (mmi *MinimumMailInput) GetPlainText() string {
	return mmi.HTMLTemplate.GetPlainText()
}

func (mmi *MinimumMailInput) GetHTMLTemplate() string {
	return mmi.HTMLTemplate.GetPlainText()
}
