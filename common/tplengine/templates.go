package tplengine

import (
	"bytes"
	"html/template"
	"io/fs"
	"io/ioutil"

	"github.com/Linus-Boehm/go-serverless-suite/entity"

	"github.com/Linus-Boehm/go-serverless-suite/itf"

	"github.com/Linus-Boehm/go-serverless-suite/common"
)

var (
	DefaultLayoutManifest = entity.TemplateManifest{
		Slot: "content",
		Path: "manifests/DefaultLayout.html",
	}

	SimpleContactFormManifest = entity.TemplateManifest{
		Slot: "content",
		Path: "manifests/mailings/en/simplecontactform.html",
	}

	CRMOptInMailManifest = entity.TemplateManifest{
		Slot: "content",
		Path: "manifests/mailings/en/CRMOptInMail.html",
	}
)

type Template struct {
	rawContent []string
	manifests  []entity.TemplateManifest
	Tpl        *template.Template
}

// LoadLayoutTemplate should be called with your root or layout template
func LoadLayoutTemplate(layout entity.TemplateManifest) (*Template, error) {
	return LoadLayoutTemplateFromFS(DefaultManifests, layout)
}

func LoadLayoutTemplateFromFS(fs fs.FS, layout entity.TemplateManifest) (*Template, error) {
	rawContent, err := open(fs, layout.Path)
	if err != nil {
		return nil, err
	}

	t := &Template{
		rawContent: []string{*rawContent},
		manifests:  []entity.TemplateManifest{layout},
	}

	t.Tpl, err = template.New("layout").Parse(*rawContent)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Template) withTemplate(fs fs.FS, manifest entity.TemplateManifest, tplName string) (*Template, error) {

	rawContent, err := open(fs, manifest.Path)
	if err != nil {
		return nil, err
	}
	t.manifests = append(t.manifests, manifest)
	t.rawContent = append(t.rawContent, *rawContent)

	_, err = t.Tpl.New(tplName).Parse(*rawContent)

	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Template) WithTemplate(fs fs.FS, manifest entity.TemplateManifest, tplName string) (itf.TplRenderer, error) {
	return t.withTemplate(fs, manifest, tplName)
}

func (t *Template) Render(data interface{}) (*string, error) {
	var buf bytes.Buffer
	err := t.Tpl.ExecuteTemplate(&buf, "layout", data)
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
