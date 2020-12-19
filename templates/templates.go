package templates

import (
	"bytes"
	"embed"
	"html/template"
	"io/ioutil"

	"github.com/markbates/pkger/pkging"

	"github.com/Linus-Boehm/go-serverless-suite/utils"
)

var (
	SimpleContactFormManifest = TemplateManifest{
		Name: "SimpleContactForm",
		Path: "manifests/mailings/en/simplecontactform.html",
	}
)

//go:embed manifests/*
var manifests embed.FS

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
	return LoadCustomTemplate(manifests, manifest)
}

func LoadCustomTemplate(fs embed.FS, manifest TemplateManifest) (*Template, error) {
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

func (t *Template) Render(data interface{}) (*string, error) {
	var buf bytes.Buffer

	err := t.Tpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return utils.StringPtr(buf.String()), nil
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
	return utils.StringPtr(string(content)), nil
}
