package handler

import (
	"goblog/internal/config"
	"net/http"
)

func NotFoundHandler(config config.BlogConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HandleError(w, r, http.StatusNotFound)
	}
}
