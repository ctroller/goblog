package block

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func RenderTemplate(templateName string, block ContentBlock) (template.HTML, error) {
	tmpl, err := template.ParseFiles(filepath.Join("ui", "templates", "blocks", templateName+".tmpl"))
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	if err := tmpl.Execute(&out, block); err != nil {
		return "", err
	}
	return template.HTML(out.String()), nil
}
