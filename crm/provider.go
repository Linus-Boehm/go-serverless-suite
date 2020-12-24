package crm

import (
	"github.com/Linus-Boehm/go-serverless-suite/templates"
)

type CRMProvider interface {
	SendSingleMail(input MinimumMailInput) error
	GetContactLists() ([]ContactList, error)
}

type Mail struct {
	Name string
	Mail string
}

type ContactList struct {
	ID             string
	Name           string
	RecipientCount int
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
	return mmi.HTMLTemplate.GetHTML()
}
