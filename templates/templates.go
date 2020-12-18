package templates

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/markbates/pkger/pkging"

	"github.com/Linus-Boehm/go-serverless-suite/utils"
	"github.com/markbates/pkger"
)

var (
	SimpleContactFormManifest = TemplateManifest{
		Name: "SimpleContactForm",
		Path: "/templates/manifests/mailings/en/simplecontactform.html",
	}
)

type TemplateManifest struct {
	Name string
	Path string
}

type Template struct {
	raw      *string
	manifest TemplateManifest
	Tpl      *template.Template
}

type TemplateOpener = func(p string) (pkging.File, error)

func LoadTemplate(manifest TemplateManifest) (*Template, error) {
	return LoadCustomTemplate(pkger.Open, manifest)
}

func LoadCustomTemplate(opener TemplateOpener, manifest TemplateManifest) (*Template, error) {
	rawContent, err := open(opener, manifest.Path)
	if err != nil {
		return nil, err
	}
	tpl, err := template.New(manifest.Name).Parse(*rawContent)
	if err != nil {
		return nil, err
	}

	return &Template{
		raw:      rawContent,
		manifest: manifest,
		Tpl:      tpl,
	}, nil
}

func (t *Template) Render(data interface{}) (*string, error) {
	var buf bytes.Buffer

	err := t.Tpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return utils.StringPtr(buf.String()), nil
}

func open(opener TemplateOpener, p string) (*string, error) {
	f, err := opener(p)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return utils.StringPtr(string(content)), nil
}
