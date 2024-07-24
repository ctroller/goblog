package handler

import (
	"goblog/internal/config"
	"goblog/internal/nav"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

// PostDetailHandler is a http.HandlerFunc that returns a post by its SEO URL
// the Route URL should be /posts/<seo-url>
func PostDetailHandler(config config.BlogConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		lastIndex := strings.LastIndex(url, "/")
		if lastIndex == -1 {
			http.Error(w, "Bad Request.", http.StatusBadRequest)
		}

		seoUrl := url[lastIndex+1:]
		post, err := config.PostService.GetBySeoURL(seoUrl)
		if err != nil || post == nil {
			HandleError(w, r, http.StatusNotFound)
			return
		}

		data := RenderData{
			Data: post,
			Breadcrumb: []nav.Breadcrumb{
				{Title: "Home", URL: "/"},
				{Title: "Posts", URL: "/posts", Nolink: true},
				{Title: post.Title, URL: "/posts/" + post.SeoURL},
			}}

		response, err := renderHTML(w, "post-detail", data)
		if err != nil {
			log.Error().Err(err).Any("post", post).Msg("Failed to render the post.")
			http.Error(w, "Failed to render post.", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
