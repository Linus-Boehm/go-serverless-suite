package tplreader

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
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
	rawContent []string
	manifests  []entity.TemplateManifest
	Tpl        *template.Template
}

func LoadTemplate(manifest entity.TemplateManifest) (*Template, error) {
	return LoadCustomTemplate(DefaultManifests, manifest)
}

func LoadCustomTemplate(fs fs.FS, manifest entity.TemplateManifest) (*Template, error) {
	temp := &Template{
		rawContent: []string{},
		manifests:  []entity.TemplateManifest{},
	}
	temp, err := temp.withTemplate(fs, manifest)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

func (t *Template) withTemplate(fs fs.FS, manifest entity.TemplateManifest) (*Template, error) {
	if t.Tpl == nil {
		t.Tpl = template.New(manifest.Name)
	}
	rawContent, err := open(fs, manifest.Path)
	if err != nil {
		return nil, err
	}
	t.manifests = append(t.manifests, manifest)
	t.rawContent = append(t.rawContent, *rawContent)
	temp, err := t.Tpl.Parse(*rawContent)
	if err != nil {
		return nil, err
	}
	t.Tpl = temp
	return t, nil
}

func (t *Template) WithTemplate(fs fs.FS, manifest entity.TemplateManifest) (itf.TplRenderer, error) {
	return t.withTemplate(fs, manifest)
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

func open(fs fs.FS, p string) (*string, error) {
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
