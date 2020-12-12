package templates

import (
	"bytes"
	"github.com/Linus-Boehm/go-serverless-suite/utils"
	"github.com/markbates/pkger"
	"html/template"
	"io/ioutil"
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
	raw *string
	manifest TemplateManifest
	Tpl *template.Template
}

func LoadTemplate(manifest TemplateManifest) (*Template,error) {
	rawContent, err := open(manifest.Path)
	if err != nil {
		return nil, err
	}
	tpl, err := template.New(manifest.Name).Parse(*rawContent)
	if err != nil {
		return nil, err
	}
	return &Template{
		raw: rawContent,
		manifest: manifest,
		Tpl: tpl,
	},nil
}

func (t *Template) Render(data interface{}) (*string,error){
	var buf bytes.Buffer

	err := t.Tpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return utils.StringPtr(string(buf.String())), nil
}


func open(p string) (*string,error){
	f, err := pkger.Open(p)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return utils.StringPtr(string(content)), nil
}