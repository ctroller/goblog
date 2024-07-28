package render

import (
	"goblog/internal/dto"
	"goblog/internal/nav"
	"goblog/internal/types"
)

func RenderPost(post *dto.Post) ([]byte, error) {
	var jsScripts = new([]types.JSScript)
	var cssFiles = new([]types.CSSFile)

	for _, block := range post.Blocks {
		if block.JSScripts() != nil {
			*jsScripts = append(*jsScripts, *block.JSScripts()...)
		}

		if block.CSSFiles() != nil {
			*cssFiles = append(*cssFiles, *block.CSSFiles()...)
		}
	}

	data := RenderData{
		Data: post,
		Breadcrumb: []nav.Breadcrumb{
			{Title: "Home", URL: "/"},
			{Title: "Posts", URL: "/posts", Nolink: true},
			{Title: post.Title, URL: "/posts/" + post.SeoURL},
		},
		JSScripts: jsScripts,
		CSSFiles:  cssFiles,
	}

	return RenderTemplate("post-detail", data)
}
