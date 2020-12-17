package mailings

import (
	"github.com/microcosm-cc/bluemonday"
)

type IMailProvider interface {
	SendSingleMail(input MinimumMailInput) error
}

type Mail struct {
	Name string
	Mail string
}

type MinimumMailInput struct {
	FromMail    Mail
	ToMail      Mail
	Subject     *string
	HTMLContent string
}

func (mmi *MinimumMailInput) GetPlainText() string {
	return bluemonday.UGCPolicy().Sanitize(mmi.HTMLContent)
}
