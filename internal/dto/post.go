package dto

import (
	"goblog/internal/block"
	"time"
)

type Post struct {
	ID         int                  `json:"id"`
	CategoryID int                  `json:"category_id"`
	Title      string               `json:"title"`
	Summary    string               `json:"summary"`
	SeoURL     string               `json:"seo_url"`
	Body       []block.ContentBlock `json:"body"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  *time.Time           `json:"updated_at,omitempty"`
	Tags       []Tag                `json:"tags"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SeoName     string `json:"seo_name"`
	Description string `json:"description"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
