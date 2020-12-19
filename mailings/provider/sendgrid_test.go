package provider

import (
	"fmt"
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/templates"

	"github.com/Linus-Boehm/go-serverless-suite/mailings"
	"github.com/Linus-Boehm/go-serverless-suite/testhelper"
	"github.com/Linus-Boehm/go-serverless-suite/utils"
	"github.com/stretchr/testify/assert"
)

func NewTestSendgridProvider(t *testing.T) *Sendgrid {
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

	validMail := mailings.Mail{
		Name: "Test Recipient",
		Mail: GetTestMailRecipient(t),
	}
	invalidMail := mailings.Mail{
		Name: "Test Recipient",
		Mail: "notexisting@unkown.qz",
	}
	tests := []struct {
		name    string
		input   mailings.MinimumMailInput
		wantErr error
	}{
		{
			name: "happy",
			input: mailings.MinimumMailInput{
				FromMail:     validMail,
				ToMail:       validMail,
				Subject:      utils.StringPtr("Test"),
				HTMLTemplate: templates.HTMLTemplate{Content: "<h1>This is a Test</h1>"},
			},
		},
		{
			name: "error not a valid email",
			input: mailings.MinimumMailInput{
				FromMail:     invalidMail,
				ToMail:       invalidMail,
				Subject:      utils.StringPtr("Test"),
				HTMLTemplate: templates.HTMLTemplate{Content: "<h1>This is a Test</h1>"},
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
