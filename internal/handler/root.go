package handler

import (
	"encoding/json"
	"goblog/internal/config"
	"goblog/internal/service"
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

		response, err := json.Marshal(posts)
		if err != nil {
			http.Error(w, "Failed to marshal posts", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
