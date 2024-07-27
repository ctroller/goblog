package block

import (
	"goblog/internal/render"
	"html/template"
)

type ContentBlock interface {
	Render() (template.HTML, error)
	DynamicScripts() *[]render.DynamicScript
	DynamicCSS() *[]render.DynamicCSS
}

type DefaultBlock struct {
	ContentBlock
	Type string
	Name string
}

func (b DefaultBlock) DynamicScripts() *[]render.DynamicScript {
	return nil
}

func (b DefaultBlock) DynamicCSS() *[]render.DynamicCSS {
	return nil
}
