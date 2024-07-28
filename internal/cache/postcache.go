package cache

import (
	"goblog/internal/config"
	"goblog/internal/dto"
	"goblog/internal/render"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func CacheAllPosts(config config.BlogConfig) {
	posts, err := config.PostService.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get cache posts.")
		return
	}

	for _, post := range posts {
		CachePost(&post, config.PostCacheConfig)
	}
}

func getPostCacheFilePath(post *dto.Post, config config.PostCacheConfig) string {
	return filepath.Join(config.CacheDir, post.SeoURL)
}

func CachePost(post *dto.Post, config config.PostCacheConfig) {
	log.Info().Any("config", config).Msg("Caching post.")
	cacheFile := getPostCacheFilePath(post, config)
	data, err := render.RenderPost(post)
	if err != nil {
		log.Error().Err(err).Int("id", post.ID).Msg("Failed to render post.")
	}

	err = os.WriteFile(cacheFile, data, 0644)
	
	if err != nil {
		log.Error().Err(err).Int("id", post.ID).Msg("Failed to save post.")
	}
}
