package mailings

import (
	"fmt"
	"os"
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/entity"

	"github.com/Linus-Boehm/go-serverless-suite/common"

	"github.com/Linus-Boehm/go-serverless-suite/testhelper"
	"github.com/stretchr/testify/assert"
)

func NewTestSendgridProvider(t *testing.T) *SendgridProvider {
	if os.Getenv("integration-test") != "1" {
		t.Skip("online test skipped")
	}
	c, err := testhelper.LoadConfig()
	assert.NoError(t, err)
	conf := SendgridConfig{
		APIKey: c.Mailings.SendgridConfig.APIKey,
	}
	return NewSendgridProvider(conf)
}

func GetTestMailRecipient(t *testing.T) string {
	c, err := testhelper.LoadConfig()
	assert.NoError(t, err)
	return c.Mailings.TestRecipient
}

func TestSendgridProvider_SendSingleMail(t *testing.T) {

	validMail := entity.Mail{
		Name: "Test Recipient",
		Mail: GetTestMailRecipient(t),
	}
	invalidMail := entity.Mail{
		Name: "Test Recipient",
		Mail: "notexisting@unkown.qz",
	}
	tests := []struct {
		name    string
		input   entity.MinimalMail
		wantErr error
	}{
		{
			name: "happy",
			input: entity.MinimalMail{
				FromMail:     validMail,
				ToMail:       validMail,
				Subject:      common.StringPtr("Test"),
				HTMLTemplate: entity.HTMLTemplate{Content: "<h1>This is a Test</h1>"},
			},
		},
		{
			name: "error not a valid email",
			input: entity.MinimalMail{
				FromMail:     invalidMail,
				ToMail:       invalidMail,
				Subject:      common.StringPtr("Test"),
				HTMLTemplate: entity.HTMLTemplate{Content: "<h1>This is a Test</h1>"},
			},
			wantErr: ErrNotAuthorizedSenderMail,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewTestSendgridProvider(t)
			fmt.Println("Sending test mail to: ", test.input.ToMail.Mail)
			err := p.SendSingleMail(test.input)
			if test.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Error(t, err)
			}
		})
	}
}

func TestSendgrid_GetContactLists(t *testing.T) {
	p := NewTestSendgridProvider(t)

	lists, err := p.GetContactLists()
	assert.NoError(t, err)
	assert.NotNil(t, lists)
}

func TestSendgridProvider_CreateUser(t *testing.T) {
	cfg, err := testhelper.LoadConfig()
	assert.NoError(t, err)

	type args struct {
		user    entity.User
		listIDs []entity.ID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy",
			args: args{
				user: entity.User{
					ID:            entity.NewEntityIDV4(),
					Email:         "test@example.org",
					Firstname:     "Max",
					Lastname:      "Muster",
					EmailVerified: true,
					Attributes: map[string]string{
						"styleInterest": "both",
					},
					Timestamps: entity.Timestamps{},
				},
				listIDs: []entity.ID{
					entity.IDFromStringOrNil(cfg.Mailings.SendgridConfig.ListID),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTestSendgridProvider(t)
			s.config.CustomFieldMap = map[string]string{
				"styleInterest": "e1_T",
			}
			if err := s.CreateUser(tt.args.user, tt.args.listIDs); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
