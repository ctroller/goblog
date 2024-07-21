package handler

import (
	"bytes"
	"encoding/json"
	"goblog/internal/config"
	"goblog/internal/service"
	"html/template"
	"log"
	"net/http"
)

func RootHandler(config config.BlogConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filter := service.Filter{
			Num: new(int),
		}

		*filter.Num = 5

		posts, err := config.PostService.GetByFilter(filter, nil, nil)
		if err != nil {
			http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
			return
		}

		var response []byte
		// check if we should render as json (get param "json" is set), or fall back to default HTML rendering
		if r.URL.Query().Has("json") {
			response, err = renderJSON(w, posts)
		} else {
			response, err = renderHTML(w, posts)
		}

		if err != nil {
			log.Fatal(err)
			http.Error(w, "Failed to render posts", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func renderJSON(w http.ResponseWriter, posts service.Posts) ([]byte, error) {
	response, err := json.Marshal(posts)
	if err != nil {
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	return response, nil
}

func renderHTML(w http.ResponseWriter, posts service.Posts) ([]byte, error) {
	tmpl, err := template.ParseFiles("ui/templates/main.html", "ui/templates/root.html")
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, posts); err != nil {
		return nil, err
	}

	w.Header().Set("Content-Type", "text/html")

	return out.Bytes(), nil
}
