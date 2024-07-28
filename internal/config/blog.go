package config

import (
	"errors"
	"goblog/internal/service"
	"os"
)

type BlogConfig struct {
	PostService     service.PostService
	PostCacheConfig PostCacheConfig
}

func (b BlogConfig) Validate() error {
	validatePostCache(b)

	return nil
}

func validatePostCache(b BlogConfig) error {
	if b.PostCacheConfig.CacheDir == "" {
		return errors.New("cache dir is required")
	}

	err := os.MkdirAll(b.PostCacheConfig.CacheDir, 0755)
	if err != nil {
		return err
	}

	return nil
}