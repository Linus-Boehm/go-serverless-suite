package mailings

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Linus-Boehm/go-serverless-suite/entity"

	"github.com/Linus-Boehm/go-serverless-suite/common"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridConfig struct {
	APIKey string
}

type Sendgrid struct {
	client *sendgrid.Client
	config *SendgridConfig
	host   string
}

var ErrNotAuthorizedSenderMail = errors.New("sender mail is not authorized")

func NewSendgridProvider(config SendgridConfig) *Sendgrid {
	c := sendgrid.NewSendClient(config.APIKey)

	return &Sendgrid{
		client: c,
		config: &config,
		host:   "https://api.sendgrid.com",
	}
}

func (s *Sendgrid) SendSingleMail(input entity.MinimalMail) error {
	plainText := input.GetPlainText()
	from := sendgridMailFromMailingsMail(input.FromMail)
	to := sendgridMailFromMailingsMail(input.ToMail)
	sbj := common.StringValue(input.Subject)
	message := mail.NewSingleEmail(from, sbj, to, plainText, input.GetHTMLTemplate())
	resp, err := s.client.Send(message)
	if err != nil {
		return err
	}
	if resp.StatusCode == 403 {
		return fmt.Errorf("%w: %s", ErrNotAuthorizedSenderMail, resp.Body)
	}
	if resp.StatusCode > 204 {
		return fmt.Errorf("unexpected response code from sendgrid: %d Msg: %s", resp.StatusCode, resp.Body)
	}
	return nil
}

func sendgridMailFromMailingsMail(m entity.Mail) *mail.Email {
	return mail.NewEmail(m.Name, m.Mail)
}

type GetContactListResponse struct {
	Result []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		ContactCount int    `json:"contact_count"`
	} `json:"result"`
}

func (s *Sendgrid) GetContactLists() ([]entity.ContactList, error) {

	request := sendgrid.GetRequest(s.config.APIKey, "/v3/marketing/lists", s.host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode > 204 {
		return nil, fmt.Errorf("unexpected response code from sendgrid: %d Msg: %s", response.StatusCode, response.Body)
	}
	payload := GetContactListResponse{}
	err = json.Unmarshal([]byte(response.Body), &payload)
	if err != nil {
		return nil, err
	}
	lists := []entity.ContactList{}

	for _, l := range payload.Result {
		lists = append(lists, entity.ContactList{
			ID:             entity.IDFromStringOrNil(l.ID),
			Name:           l.Name,
			RecipientCount: l.ContactCount,
		})
	}
	return lists, nil
}
