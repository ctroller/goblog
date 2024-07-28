package render

import (
	"bytes"
	"goblog/internal/nav"
	"goblog/internal/types"
	"html/template"
	"net/http"
)

type RenderData struct {
	Data       any
	Breadcrumb []nav.Breadcrumb
	JSScripts  *[]types.JSScript
	CSSFiles   *[]types.CSSFile
}

type TemplateRenderData struct {
	Data            any
	Breadcrumb      []nav.Breadcrumb
	HeaderJSScripts *[]types.JSScript
	BodyJSScripts   *[]types.JSScript
	CSSFiles        *[]types.CSSFile
}

// renders given template (within the ui/templates directory) with the provided data.
// the renderer will convert the given data to a struct with a Data field (containing the provided data) and a Referrer field (containing the http referrer if provided).
func RenderHTML(w http.ResponseWriter, templateName string, data RenderData) ([]byte, error) {
	out, err := RenderTemplate(templateName, data)
	w.Header().Set("Content-Type", "text/html")

	return out, err
}

func RenderTemplate(templateName string, data RenderData) ([]byte, error) {
	tmpl, err := template.ParseFiles("ui/templates/main.tmpl", "ui/templates/nav/breadcrumb.tmpl", "ui/templates/"+templateName+".tmpl")
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, toTemplateData(data)); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func toTemplateData(data RenderData) TemplateRenderData {
	templateData := TemplateRenderData{
		Data:       data.Data,
		Breadcrumb: data.Breadcrumb,
	}

	if data.CSSFiles != nil {
		templateData.CSSFiles = getUniqueCSS(*data.CSSFiles)
	}

	if data.JSScripts != nil {
		templateData.HeaderJSScripts = getUniqueScripts(*data.JSScripts, false)
		templateData.BodyJSScripts = getUniqueScripts(*data.JSScripts, true)
	}

	return templateData
}

func getUniqueCSS(dynamicCSS []types.CSSFile) *[]types.CSSFile {
	uniqueCSSLinks := make(map[string]bool)
	var uniqueCSS []types.CSSFile

	for _, css := range dynamicCSS {
		if !uniqueCSSLinks[css.Link] {
			uniqueCSSLinks[css.Link] = true
			uniqueCSS = append(uniqueCSS, css)
		}
	}

	return &uniqueCSS
}

func getUniqueScripts(dynamicScripts []types.JSScript, footer bool) *[]types.JSScript {
	uniqueScriptLinks := make(map[string]bool)
	uniqueScripts := make(map[string]bool)
	var scripts []types.JSScript

	for _, script := range dynamicScripts {
		if script.Footer != footer {
			continue
		}
		if script.Link != "" {
			if !uniqueScriptLinks[script.Link] {
				uniqueScriptLinks[script.Link] = true
				scripts = append(scripts, script)
			}
		} else {
			if script.UniqueId == "" || !uniqueScripts[script.UniqueId] {
				uniqueScripts[script.UniqueId] = true
				scripts = append(scripts, script)
			}
		}
	}

	return &scripts
}
