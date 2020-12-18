package provider

import (
	"errors"
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/mailings"
	"github.com/Linus-Boehm/go-serverless-suite/utils"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridConfig struct {
	APIKey string
}

type Sendgrid struct {
	client *sendgrid.Client
}

var ErrNotAuthorizedSenderMail = errors.New("sender mail is not authorized")

func NewSendgridProvider(config SendgridConfig) *Sendgrid {
	c := sendgrid.NewSendClient(config.APIKey)

	return &Sendgrid{
		client: c,
	}
}

func (s *Sendgrid) SendSingleMail(input mailings.MinimumMailInput) error {
	plainText := input.GetPlainText()
	from := sendgridMailFromMailingsMail(input.FromMail)
	to := sendgridMailFromMailingsMail(input.ToMail)
	sbj := utils.StringValue(input.Subject)
	message := mail.NewSingleEmail(from, sbj, to, plainText, input.HTMLContent)
	resp, err := s.client.Send(message)
	if err != nil {
		return err
	}
	var errorMsg string
	if resp != nil {
		errorMsg = resp.Body
	}
	if resp.StatusCode == 403 {
		return fmt.Errorf("%w: %s", ErrNotAuthorizedSenderMail, errorMsg)
	}
	if resp.StatusCode > 204 {
		return fmt.Errorf("unexpected response code from sendgrid: %d Msg: %s", resp.StatusCode, errorMsg)
	}
	return nil
}

func sendgridMailFromMailingsMail(m mailings.Mail) *mail.Email {
	return mail.NewEmail(m.Name, m.Mail)
}
