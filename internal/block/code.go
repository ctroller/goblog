package block

import (
	"goblog/internal/render"
	"html/template"
)

const hJSVersion = "11.10.0"

var supportedHJSLanguages = []string{
	"bash", "css", "dockerfile", "go", "http", "ini", "java", "javascript", "json", "kotlin", "lua", "markdown", "nginx", "nginx", "pgsql", "protobuf", "python", "scss", "sql", "typescript", "xml", "yaml",
}

const hJSCDNLink = "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/" + hJSVersion
const hJSCDNLinkPrefix = hJSCDNLink + "/languages/"
const hJSCSSLink = hJSCDNLink + "/styles/base16/gruvbox-dark-soft.min.css"

var supportedHJSLanguagesWithPrefix = func() []string {
	prefixedLanguages := make([]string, len(supportedHJSLanguages))
	for i, lang := range supportedHJSLanguages {
		prefixedLanguages[i] = hJSCDNLinkPrefix + lang + ".min.js"
	}
	return prefixedLanguages
}()

var dynamicScripts = func() *[]render.DynamicScript {
	var scripts []render.DynamicScript
	scripts = append(scripts, render.DynamicScript{
		Link: hJSCDNLink + "/highlight.min.js", Async: false, Footer: true,
	})

	for _, lang := range supportedHJSLanguagesWithPrefix {
		scripts = append(scripts, render.DynamicScript{
			Link:   lang,
			Async:  true,
			Footer: true,
		})
	}

	scripts = append(scripts, render.DynamicScript{
		Link: "/static/js/hljs-copy.js",
		Async: true,
		Footer: true,
	})

	scripts = append(scripts, render.DynamicScript{
		Content: `hljs.highlightAll();`,
		Async:   true,
		Footer:  true,
		UniqueId: "highlightjs-init",
	})

	return &scripts
}

type CodeBlock struct {
	DefaultBlock
	Content  string
	Language string
}

func (b CodeBlock) Render() (template.HTML, error) {
	return RenderTemplate("Code", b)
}

func (b CodeBlock) DynamicScripts() *[]render.DynamicScript {
	return dynamicScripts()
}

func (b CodeBlock) DynamicCSS() *[]render.DynamicCSS {
	return &[]render.DynamicCSS{
		{Link: hJSCSSLink},
		{Link: "https://unpkg.com/highlightjs-copy/dist/highlightjs-copy.min.css"},
	}
}

func NewCodeBlock(content string, language string) CodeBlock {
	return CodeBlock{
		DefaultBlock: DefaultBlock{
			Type: "code",
			Name: "Code",
		},
		Content:  content,
		Language: language,
	}
}
