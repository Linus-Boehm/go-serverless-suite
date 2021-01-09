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
	APIKey        string
	DefaultSender *entity.Mail
}

type SendgridProvider struct {
	client *sendgrid.Client
	config *SendgridConfig
	host   string
}

var ErrNotAuthorizedSenderMail = errors.New("sender mail is not authorized")

func NewSendgridProvider(config SendgridConfig) *SendgridProvider {
	c := sendgrid.NewSendClient(config.APIKey)

	return &SendgridProvider{
		client: c,
		config: &config,
		host:   "https://api.sendgrid.com",
	}
}

func (s *SendgridProvider) SendSingleMail(input entity.MinimalMail) error {
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

func (s *SendgridProvider) GetDefaultSender() *entity.Mail {
	return s.config.DefaultSender
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

func (s *SendgridProvider) GetContactLists() ([]entity.MailContactList, error) {
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
	lists := []entity.MailContactList{}

	for _, l := range payload.Result {
		lists = append(lists, entity.MailContactList{
			ID:             entity.IDFromStringOrNil(l.ID),
			Name:           l.Name,
			RecipientCount: l.ContactCount,
		})
	}
	return lists, nil
}

type SendgridContact struct {
	Email        string                 `json:"email"`
	FirstName    string                 `json:"first_name"`
	LastName     string                 `json:"last_name"`
	CustomFields map[string]interface{} `json:"custom_fields"`
}

type SendgridUpdateContactRequest struct {
	ListIDS  []string          `json:"list_ids"`
	Contacts []SendgridContact `json:"contacts"`
}

func mapUserEntityToContact(user entity.User) SendgridContact {
	attrs := map[string]interface{}{}
	for k, v := range user.Attributes {
		attrs[k] = v
	}
	return SendgridContact{
		Email:        user.Email,
		FirstName:    user.Firstname,
		LastName:     user.Lastname,
		CustomFields: attrs,
	}
}

// listIDs of the sendgrid lists, if null
func (s *SendgridProvider) CreateUser(user entity.User, listIDs []entity.ID) error {
	listPayload := []string{}
	for _, id := range listIDs {
		listPayload = append(listPayload, id.String())
	}
	requestPayload := SendgridUpdateContactRequest{
		ListIDS: listPayload,
		Contacts: []SendgridContact{
			mapUserEntityToContact(user),
		},
	}
	jsonPayload, err := json.Marshal(requestPayload)
	if err != nil {
		return err
	}

	request := sendgrid.GetRequest(s.config.APIKey, "/v3/marketing/contacts", s.host)
	request.Method = "PUT"
	request.Body = jsonPayload
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	}
	if response.StatusCode == 403 {
		return fmt.Errorf("%w: %s", ErrNotAuthorizedSenderMail, response.Body)
	}
	if response.StatusCode != 202 {
		return fmt.Errorf("unexpected response code from sendgrid: %d Msg: %s", response.StatusCode, response.Body)
	}
	return nil
}
