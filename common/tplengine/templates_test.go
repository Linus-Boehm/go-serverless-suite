package tplengine

import (
	"bytes"
	"html/template"
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
					"Slot": "Max Muster",
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
			tpl, err := LoadLayoutTemplate(test.input)
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
		Slot: "ExampleLayout",
		Path: "manifests/examples/ExampleLayout.html",
	}
	contentManifest := entity.TemplateManifest{
		Slot: "ExampleContent",
		Path: "manifests/examples/ExampleContent.html",
	}
	footerManifest := entity.TemplateManifest{
		Slot: "ExampleFooter",
		Path: "manifests/examples/ExampleFooter.html",
	}
	type input struct {
		template entity.TemplateManifest
		slot     string
	}
	type data struct {
		Title    string
		Subtitle string
	}
	tests := []struct {
		name              string
		inputs            []input
		data              data
		wantRenderErr     bool
		mustIncludeRender []string
	}{
		{
			name: "happy",
			inputs: []input{
				{
					template: layoutManifest,
					slot:     "layout",
				},
				{
					template: contentManifest,
					slot:     "content",
				},
				{
					template: footerManifest,
					slot:     "footer",
				},
			},
			data: data{
				Title:    "TestTitle",
				Subtitle: "TestSubtitle",
			},
			mustIncludeRender: []string{"TestTitle", "TestSubtitle", "Thank you for templating"},
		},
		{
			name: "error no footer",
			inputs: []input{
				{
					template: layoutManifest,
					slot:     "layout",
				},
				{
					template: contentManifest,
					slot:     "content",
				},
			},
			data: data{
				Title:    "TestTitle",
				Subtitle: "TestSubtitle",
			},
			mustIncludeRender: []string{"TestTitle", "TestSubtitle"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tpl, err := LoadLayoutTemplate(test.inputs[0].template)
			assert.NoError(t, err)
			assert.NotNil(t, tpl)
			if len(test.inputs) > 1 {
				for i := 1; i < len(test.inputs); i++ {
					_, err := tpl.WithTemplate(DefaultManifests, test.inputs[i].template, test.inputs[i].slot)
					assert.NoError(t, err)
				}
			}
			result, err := tpl.RenderWithHTML(test.data)
			if !test.wantRenderErr {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if test.wantRenderErr != (err != nil) {
					t.Errorf("RenderWithHTML() error = %v", err)
				}

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

func Test_templateing(t *testing.T) {
	type data struct {
		Title    string
		Subtitle string
	}
	input := data{
		Title:    "Title",
		Subtitle: "Subtitle",
	}
	rawContent, err := open(DefaultManifests, "manifests/examples/ExamplePlainContent.html")
	assert.NoError(t, err)
	rawLayout, err := open(DefaultManifests, "manifests/examples/ExampleLayout.html")
	assert.NoError(t, err)
	tpl := template.Must(template.New("layout").Parse(*rawLayout))
	tpl.New("content").Parse(*rawContent)
	tpl.New("footer").Parse(*rawContent)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, input)
	assert.NoError(t, err)
	result := buf.String()
	assert.Contains(t, result, input.Subtitle)
}
