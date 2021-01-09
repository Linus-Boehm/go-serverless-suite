package services

import (
	"strings"
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/common/tplengine"
	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
	"github.com/golang/mock/gomock"
)

func Test_crmSVC_SendDoubleOptInMail(t *testing.T) {

	type args struct {
		options entity.CRMOptInMailOptions
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		expectContains []string
	}{
		{
			name: "happy",
			args: args{
				options: entity.CRMOptInMailOptions{
					Fullname:         "Max Mustermann",
					Subject:          nil,
					EMail:            entity.IDFromStringOrNil("test@example.org"),
					UserID:           entity.IDFromStringOrNil("123"),
					SubID:            entity.IDFromStringOrNil("456"),
					ConfirmationPath: "http://localhost",
					Template:         &tplengine.CRMOptInMailManifest,
					FS:               &tplengine.DefaultManifests,
				},
			},
			wantErr: false,
			expectContains: []string{
				"Max Mustermann",
				"http://localhost/?id=123&subid=456&email=test%40example.org",
			},
		},
		{
			name: "happy empty",
			args: args{
				options: entity.CRMOptInMailOptions{
					Fullname:         "Max Mustermann",
					Subject:          nil,
					EMail:            entity.IDFromStringOrNil("test@example.org"),
					UserID:           entity.IDFromStringOrNil("123"),
					SubID:            entity.IDFromStringOrNil("456"),
					ConfirmationPath: "http://localhost",
				},
			},
			wantErr: false,
			expectContains: []string{
				"Max Mustermann",
				"http://localhost/?id=123&subid=456&email=test%40example.org",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mailProvider := itf.NewMockMailerProvider(ctrl)
			mailer := itf.NewMockMailer(ctrl)

			mailer.EXPECT().GetProvider().Return(mailProvider)
			c := NewCRMService(mailer, nil, nil, entity.Mail{
				Mail: "test@example.org",
			})
			mailProvider.EXPECT().SendSingleMail(gomock.Any()).DoAndReturn(func(mail entity.MinimalMail) error {
				html := mail.HTMLTemplate.GetHTML()
				for _, input := range tt.expectContains {
					if !strings.Contains(html, input) {
						t.Errorf("SendDoubleOptInMail() expectStr = %s, in outpur %s", input, html)
					}
				}
				return nil
			})
			if err := c.SendDoubleOptInMail(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("SendDoubleOptInMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
