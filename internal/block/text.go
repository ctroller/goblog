package block

import "html/template"

type TextBlock struct {
	DefaultBlock
	Content string `json:"content"`
}

func (b TextBlock) Render() (template.HTML, error) {
	return RenderTemplate("Text", b)
}

func NewTextBlock(content string) TextBlock {
	return TextBlock{
		DefaultBlock: DefaultBlock{
			Type: "text",
			Name: "Text",
		},
		Content: content,
	}
}
