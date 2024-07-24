package handler

import (
	"goblog/internal/config"
	"goblog/internal/nav"
	"goblog/internal/service"
	"net/http"

	"github.com/rs/zerolog/log"
)

func RootHandler(config config.BlogConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filter := service.Filter{
			Num: new(int),
		}

		*filter.Num = 5

		posts, err := config.PostService.GetByFilter(filter, nil, nil)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to fetch posts")
		}

		data := RenderData{
			Data:       posts,
			Breadcrumb: []nav.Breadcrumb{{Title: "Home", URL: "/"}},
		}

		response, err := renderHTML(w, "root", data)

		if err != nil {
			log.Error().Err(err).Msg("Failed to render posts")
			http.Error(w, "Failed to render posts.", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
