package mailer

import (
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMailingsService_SendSimpleContactForm(t *testing.T) {
	tests := []struct {
		name              string
		input             entity.ContactForm
		expectRenderCalls func(recorder *itf.MockTplRendererMockRecorder)
		expectMailerCalls func(recorder *itf.MockMailerProviderMockRecorder)
		expectErr         bool
	}{
		{
			name: "happy",
			input: entity.ContactForm{
				SenderMail: entity.Mail{
					Name: "sender",
					Mail: "SenderMail",
				},
				FromMail: entity.Mail{
					Name: "from",
					Mail: "fromMail",
				},
				ToMail: entity.Mail{
					Name: "to",
					Mail: "toMail",
				},
				Subject: "Sub",
				Message: "MSG",
				Attributes: map[string]string{
					"test": "T",
				},
			},
			expectRenderCalls: func(recorder *itf.MockTplRendererMockRecorder) {
				recorder.RenderWithHTML(gomock.Any()).Return(&entity.HTMLTemplate{}, nil)
			},
			expectMailerCalls: func(recorder *itf.MockMailerProviderMockRecorder) {
				recorder.SendSingleMail(gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRender := itf.NewMockTplRenderer(ctrl)
			mockMailProvider := itf.NewMockMailerProvider(ctrl)

			mailer := NewMailerService(mockMailProvider)

			if test.expectRenderCalls != nil {
				test.expectRenderCalls(mockRender.EXPECT())
			}

			if test.expectMailerCalls != nil {
				test.expectMailerCalls(mockMailProvider.EXPECT())
			}

			err := mailer.SendSimpleContactForm(test.input, mockRender)
			assert.Equal(t, test.expectErr, err != nil)
		})
	}
}
