package tplreader

import (
	"testing"

	"github.com/Linus-Boehm/go-serverless-suite/entity"

	"github.com/stretchr/testify/assert"
)

func TestLoadRenderTemplate(t *testing.T) {
	tests := []struct {
		name              string
		input             entity.TemplateManifest
		data              interface{}
		wantRenderErr     bool
		mustIncludeRender []string
	}{
		{
			name:  "happy",
			input: SimpleContactFormManifest,
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
			mustIncludeRender: []string{"TestValue", "New Contact Request"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tpl, err := LoadTemplate(test.input)
			assert.NoError(t, err)
			assert.NotNil(t, tpl)
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

func TestLoadRenderTemplateWithLayouts(t *testing.T) {
	layoutManifest := entity.TemplateManifest{
		Name: "ExampleLayout",
		Path: "manifests/examples/ExampleLayout.html",
	}
	contentManifest := entity.TemplateManifest{
		Name: "ExampleContent",
		Path: "manifests/examples/ExampleContent.html",
	}
	type data struct {
		Title    string
		Subtitle string
	}
	tests := []struct {
		name              string
		inputs            []entity.TemplateManifest
		data              data
		wantRenderErr     bool
		mustIncludeRender []string
	}{
		{
			name:   "happy",
			inputs: []entity.TemplateManifest{contentManifest, layoutManifest},
			data: data{
				Title:    "TestTitle",
				Subtitle: "TestSubtitle",
			},
			mustIncludeRender: []string{"TestTitle", "TestSubtitle"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tpl, err := LoadTemplate(test.inputs[0])
			assert.NoError(t, err)
			assert.NotNil(t, tpl)
			if len(test.inputs) > 1 {
				for i := 1; i < len(test.inputs); i++ {
					_, err := tpl.WithTemplate(DefaultManifests, test.inputs[i])
					assert.NoError(t, err)
				}
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
