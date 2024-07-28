package cache

import (
	"goblog/internal/config"
	"goblog/internal/dto"
	"goblog/internal/handler"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const cacheDir = "cache/posts"

func CacheAllPosts(config config.BlogConfig) {
	posts, err := config.PostService.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get cache posts.")
		return
	}

	for _, post := range posts {
		CachePost(&post)
	}
}

func getCacheFilePath(post *dto.Post) string {
	return filepath.Join(cacheDir, post.SeoURL)
}

func CachePost(post *dto.Post) error {
	cacheFile := getCacheFilePath(post)
	data, err := handler.RenderPost(post)
	if err != nil {
		return err
	}

	return os.WriteFile(cacheFile, data, 0644)
}
