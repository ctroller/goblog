package config

import (
	"goblog/internal/service"
)

type BlogConfig struct {
	PostService service.PostService
}
