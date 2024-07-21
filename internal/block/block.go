package block

import "html/template"

type ContentBlock interface {
	Render() (template.HTML, error)
}

type DefaultBlock struct {
	ContentBlock `json:"-"`
	Type         string `json:"type"`
	Name         string `json:"-"`
}
