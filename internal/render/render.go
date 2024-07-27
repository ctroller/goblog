package render

import (
	"bytes"
	"goblog/internal/nav"
	"html/template"
	"net/http"
)

type DynamicScript struct {
	Link     string
	Content  template.JS
	Async    bool
	Footer   bool
	UniqueId string
}

type DynamicCSS struct {
	Link string
}

type RenderData struct {
	Data           interface{}
	Breadcrumb     []nav.Breadcrumb
	DynamicScripts *[]DynamicScript
	DynamicCSS     *[]DynamicCSS
}

type TemplateRenderData struct {
	Data               interface{}
	Breadcrumb         []nav.Breadcrumb
	DynamicHeadScripts *[]DynamicScript
	DynamicBodyScripts *[]DynamicScript
	DynamicCSS         *[]DynamicCSS
}

// renders given template (within the ui/templates directory) with the provided data.
// the renderer will convert the given data to a struct with a Data field (containing the provided data) and a Referrer field (containing the http referrer if provided).
func RenderHTML(w http.ResponseWriter, templateName string, data RenderData) ([]byte, error) {
	tmpl, err := template.ParseFiles("ui/templates/main.tmpl", "ui/templates/nav/breadcrumb.tmpl", "ui/templates/"+templateName+".tmpl")
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, toTemplateData(data)); err != nil {
		return nil, err
	}

	w.Header().Set("Content-Type", "text/html")

	return out.Bytes(), nil
}

func toTemplateData(data RenderData) TemplateRenderData {
	templateData := TemplateRenderData{
		Data:       data.Data,
		Breadcrumb: data.Breadcrumb,
	}

	if data.DynamicCSS != nil {
		templateData.DynamicCSS = getUniqueCSS(*data.DynamicCSS)
	}

	if data.DynamicScripts != nil {
		templateData.DynamicHeadScripts = getUniqueScripts(*data.DynamicScripts, false)
		templateData.DynamicBodyScripts = getUniqueScripts(*data.DynamicScripts, true)
	}

	return templateData
}

func getUniqueCSS(dynamicCSS []DynamicCSS) *[]DynamicCSS {
	uniqueCSSLinks := make(map[string]bool)
	var uniqueCSS []DynamicCSS

	for _, css := range dynamicCSS {
		if !uniqueCSSLinks[css.Link] {
			uniqueCSSLinks[css.Link] = true
			uniqueCSS = append(uniqueCSS, css)
		}
	}

	return &uniqueCSS
}

func getUniqueScripts(dynamicScripts []DynamicScript, footer bool) *[]DynamicScript {
	uniqueScriptLinks := make(map[string]bool)
	uniqueScripts := make(map[string]bool)
	var scripts []DynamicScript

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
