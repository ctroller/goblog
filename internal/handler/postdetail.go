package handler

import (
	"goblog/internal/config"
	"goblog/internal/dto"
	"goblog/internal/nav"
	"goblog/internal/render"
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

		response, err := RenderPost(post)
		if err != nil {
			log.Error().Err(err).Any("post", post).Msg("Failed to render the post.")
			http.Error(w, "Failed to render post.", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func RenderPost(post *dto.Post) ([]byte, error) {
	var dynamicScripts = new([]render.DynamicScript)
	var dynamicCSS = new([]render.DynamicCSS)
	for _, block := range post.Body {
		if block.DynamicScripts() != nil {
			*dynamicScripts = append(*dynamicScripts, *block.DynamicScripts()...)
		}

		if block.DynamicCSS() != nil {
			*dynamicCSS = append(*dynamicCSS, *block.DynamicCSS()...)
		}
	}

	data := render.RenderData{
		Data: post,
		Breadcrumb: []nav.Breadcrumb{
			{Title: "Home", URL: "/"},
			{Title: "Posts", URL: "/posts", Nolink: true},
			{Title: post.Title, URL: "/posts/" + post.SeoURL},
		},
		DynamicScripts: dynamicScripts,
		DynamicCSS:     dynamicCSS,
	}

	return render.RenderTemplate("post-detail", data)
}
