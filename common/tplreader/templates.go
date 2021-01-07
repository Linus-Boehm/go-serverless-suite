package tplreader

import (
	"bytes"
	"embed"
	"html/template"
	"io/ioutil"

	"github.com/Linus-Boehm/go-serverless-suite/entity"

	"github.com/Linus-Boehm/go-serverless-suite/itf"

	"github.com/Linus-Boehm/go-serverless-suite/common"
)

var (
	SimpleContactFormManifest = entity.TemplateManifest{
		Name: "SimpleContactForm",
		Path: "manifests/mailings/en/simplecontactform.html",
	}

	CRMOptInMailManifest = entity.TemplateManifest{
		Name: "CRMOptInMail",
		Path: "manifests/mailings/en/CRMOptInMail.html",
	}
)

//go:embed manifests/*
var DefaultManifests embed.FS

type Template struct {
	raw      *string
	manifest entity.TemplateManifest
	Tpl      *template.Template
}

func LoadTemplate(manifest entity.TemplateManifest) (itf.TplRenderer, error) {
	return LoadCustomTemplate(DefaultManifests, manifest)
}

func LoadCustomTemplate(fs embed.FS, manifest entity.TemplateManifest) (itf.TplRenderer, error) {
	rawContent, err := open(fs, manifest.Path)
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

func (t *Template) GetRaw() *string {
	return t.raw
}

func (t *Template) Render(data interface{}) (*string, error) {
	var buf bytes.Buffer
	err := t.Tpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return common.StringPtr(buf.String()), nil
}

func (t *Template) RenderWithHTML(data interface{}) (*entity.HTMLTemplate, error) {
	content, err := t.Render(data)
	if err != nil {
		return nil, err
	}
	return &entity.HTMLTemplate{Content: *content}, nil
}

func open(fs embed.FS, p string) (*string, error) {
	f, err := fs.Open(p)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return common.StringPtr(string(content)), nil
}
