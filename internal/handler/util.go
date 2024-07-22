package handler

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

func renderHTML(w http.ResponseWriter, templateName string, data interface{}) ([]byte, error) {
	tmpl, err := template.ParseFiles("ui/templates/main.html", "ui/templates/"+templateName+".html")
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
	response, err := renderHTML(w, "error/"+strconv.Itoa(status), nil)
	if err != nil {
		log.Error().Err(err).Int("status_code", status).Msg("Failed to render error page.")
		http.Error(w, "Failed to render error page.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(response)
}
