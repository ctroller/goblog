package block

import (
	"goblog/internal/render"
	"html/template"
)

var supportedHJSLanguages = []string{
	"bash", "css", "dockerfile", "go", "http", "ini", "java", "javascript", "json", "kotlin", "lua", "markdown", "nginx", "nginx", "pgsql", "protobuf", "python", "scss", "sql", "typescript", "xml", "yaml",
}

const hJSVersion = "11.10.0"
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

func createDynamicScript(link string, async bool, footer bool, content template.JS, uniqueId string) render.DynamicScript {
	return render.DynamicScript{
		Link:     link,
		Async:    async,
		Footer:   footer,
		Content:  content,
		UniqueId: uniqueId,
	}
}

var dynamicScripts = func() *[]render.DynamicScript {
	var scripts []render.DynamicScript

	scripts = append(scripts, createDynamicScript(hJSCDNLink+"/highlight.min.js", false, true, "", ""))
	for _, lang := range supportedHJSLanguagesWithPrefix {
		scripts = append(scripts, createDynamicScript(lang, true, true, "", ""))
	}

	scripts = append(scripts, createDynamicScript("/static/js/hljs-copy.js", true, true, "", ""))
	scripts = append(scripts, createDynamicScript("", true, true, `hljs.highlightAll();`, "highlightjs-init"))

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
