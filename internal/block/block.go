package block

import (
	"goblog/internal/types"
	"html/template"
)

type ContentBlock interface {
	Render() (template.HTML, error)
	JSScripts() *[]types.JSScript
	CSSFiles() *[]types.CSSFile
}

type DefaultBlock struct {
	ContentBlock
	Type string
	Name string
}

func (b DefaultBlock) JSScripts() *[]types.JSScript {
	return nil
}

func (b DefaultBlock) CSSFiles() *[]types.CSSFile {
	return nil
}
