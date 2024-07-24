package handler

import (
	"bytes"
	"goblog/internal/nav"
	"html/template"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

type RenderData struct {
	Data       interface{}
	Breadcrumb []nav.Breadcrumb
}

// renders given template (within the ui/templates directory) with the provided data.
// the renderer will convert the given data to a struct with a Data field (containing the provided data) and a Referrer field (containing the http referrer if provided).
func renderHTML(w http.ResponseWriter, templateName string, data RenderData) ([]byte, error) {
	tmpl, err := template.ParseFiles("ui/templates/main.html", "ui/templates/nav/breadcrumb.html", "ui/templates/"+templateName+".html")
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, data); err != nil {
		return nil, err
	}

	w.Header().Set("Content-Type", "text/html")

	return out.Bytes(), nil
}

func HandleError(w http.ResponseWriter, r *http.Request, status int) {
	data := RenderData{
		Data:       nil,
		Breadcrumb: []nav.Breadcrumb{{Title: "Home", URL: "/"}}}
	response, err := renderHTML(w, "error/"+strconv.Itoa(status), data)
	if err != nil {
		log.Error().Err(err).Int("status_code", status).Msg("Failed to render error page.")
		http.Error(w, "Failed to render error page.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(response)
}
