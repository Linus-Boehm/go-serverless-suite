package provider

import (
	"fmt"
	"github.com/Linus-Boehm/go-serverless-suite/mailings"
	"github.com/Linus-Boehm/go-serverless-suite/testhelper"
	"github.com/Linus-Boehm/go-serverless-suite/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewTestSendgridProvider(t *testing.T) *SendgridProvider {
	c, err := testhelper.LoadConfig()
	assert.NoError(t, err)
	conf := SendgridConfig{
		APIKey: c.Mailings.SendgridConfig.ApiKey,
	}
	return NewSendgridProvider(conf)
}

func GetTestMailRecipient(t *testing.T) string {
	c, err := testhelper.LoadConfig()
	assert.NoError(t, err)
	return c.Mailings.TestRecipient
}

func TestSendgridProvider_SendSingleMail(t *testing.T) {

	toMail := mailings.Mail{
		Name: "Test Recipient",
		Mail: GetTestMailRecipient(t),
	}
	tests := []struct {
		name string
		input mailings.MinimumMailInput
	}{
		{
			name: "happy",
			input: mailings.MinimumMailInput{
				FromMail:    toMail,
				ToMail:      toMail,
				Subject:     utils.StringPtr("Test"),
				HtmlContent: "<h1>This is a Test</h1>",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewTestSendgridProvider(t)
			fmt.Println("Sending test mail to: ", toMail.Mail)
			err := p.SendSingleMail(test.input)
			assert.NoError(t, err)
		})
	}
}
