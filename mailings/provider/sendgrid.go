package provider

import (
	"errors"
	"fmt"
	"github.com/Linus-Boehm/go-serverless-suite/mailings"
	"github.com/Linus-Boehm/go-serverless-suite/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridConfig struct {
	APIKey string
}

type SendgridProvider struct {
	client *sendgrid.Client
}

var NotAuthorizedSenderMail = errors.New("sender mail is not authorized")

func NewSendgridProvider(config SendgridConfig) *SendgridProvider {
	c := sendgrid.NewSendClient(config.APIKey)

	return &SendgridProvider{
		client: c,
	}
}


func (s *SendgridProvider) SendSingleMail(input mailings.MinimumMailInput) error {
	plainText := input.GetPlainText()
	from := sendgridMailFromMailingsMail(input.FromMail)
	to := sendgridMailFromMailingsMail(input.ToMail)
	sbj:= utils.StringValue(input.Subject)
	message := mail.NewSingleEmail(from, sbj, to, plainText, input.HtmlContent)
	resp, err := s.client.Send(message)
	if err != nil {
		return err
	}
	if resp.StatusCode == 403 {
		return NotAuthorizedSenderMail
	}
	if resp.StatusCode > 204 {
		return fmt.Errorf("unexpected response code from sendgrid: %d", resp.StatusCode)
	}
	return nil
}

func sendgridMailFromMailingsMail(m mailings.Mail) *mail.Email {
	return mail.NewEmail(m.Name, m.Mail)
}