package templates

import (
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/common"

	"github.com/stretchr/testify/assert"
)

func TestLoadRenderTemplate(t *testing.T) {
	tests := []struct {
		name              string
		input             TemplateManifest
		data              interface{}
		mustInclude       *string
		wantRenderErr     bool
		mustIncludeRender []string
	}{
		{
			name:        "happy",
			input:       SimpleContactFormManifest,
			mustInclude: common.StringPtr("New Contact Request"),
			data: map[string]interface{}{
				"FromMail": map[string]interface{}{
					"Name": "Max Muster",
					"Mail": "max.muster@example.org",
				},
				"Subject": "Subject",
				"Message": "Message",
				"Attributes": map[string]string{
					"TestAttr": "TestValue",
				},
			},
			mustIncludeRender: []string{"TestValue"},
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
			result, err := tpl.RenderWithHTML(test.data)
			if !test.wantRenderErr {
				assert.NoError(t, err)
				assert.NotNil(t, result)

			} else {
				assert.Nil(t, err)
			}
			if test.mustIncludeRender != nil {
				for _, incl := range test.mustIncludeRender {
					assert.Contains(t, result.GetHTML(), incl)
				}

			}
		})
	}
}
