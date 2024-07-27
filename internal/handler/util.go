package handler

import (
	"goblog/internal/nav"
	"goblog/internal/render"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

func HandleError(w http.ResponseWriter, r *http.Request, status int) {
	data := render.RenderData{
		Data:       nil,
		Breadcrumb: []nav.Breadcrumb{{Title: "Home", URL: "/"}}}
	response, err := render.RenderHTML(w, "error/"+strconv.Itoa(status), data)
	if err != nil {
		log.Error().Err(err).Int("status_code", status).Msg("Failed to render error page.")
		http.Error(w, "Failed to render error page.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(response)
}
