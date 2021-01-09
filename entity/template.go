package entity

import "github.com/microcosm-cc/bluemonday"

type HTMLTemplate struct {
	Title       string
	Content     string
	TextContent string
}

type TemplateManifest struct {
	Name string
	Path string
}

func (h *HTMLTemplate) GetPlainText() string {
	return bluemonday.StripTagsPolicy().Sanitize(h.Content)
}

func (h *HTMLTemplate) GetHTML() string {
	return h.Content
}
