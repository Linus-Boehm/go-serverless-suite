package templates

import (
	"github.com/Linus-Boehm/go-serverless-suite/mailings"
	"github.com/Linus-Boehm/go-serverless-suite/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadRenderTemplate(t *testing.T) {
	tests := []struct {
		name string
		input TemplateManifest
		data interface{}
		mustInclude *string
		wantRenderErr bool
		mustIncludeRender *string
	}{
		{
			name: "happy",
			input: SimpleContactFormManifest,
			mustInclude: utils.StringPtr("New Contact Request"),
			data: mailings.SimpleContactFormInput{
				FromMail:       mailings.Mail{
					Name: "Max Muster",
					Mail: "max.muster@example.org",
				},
				Subject:        "Subject",
				Message:        "Message",
				Attributes: map[string]string{
					"TestAttr": "TestValue",
				},
			},
			mustIncludeRender: utils.StringPtr("TestValue"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tpl, err := LoadTemplate(test.input)
			assert.NoError(t, err)
			assert.NotNil(t, tpl)
			if test.mustInclude != nil {
				assert.Contains(t, *tpl.raw, *test.mustInclude)
			}
			result, err := tpl.Render(test.data)
			if !test.wantRenderErr{
				assert.NoError(t, err)
				assert.NotNil(t, result)

			}else{
				assert.Nil(t, err)
			}
			if test.mustIncludeRender != nil {
				assert.Contains(t, *result, *test.mustIncludeRender)
			}
		})
	}
}


