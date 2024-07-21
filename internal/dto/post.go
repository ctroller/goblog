package dto

import "time"

type Post struct {
	ID         int        `json:"id"`
	CategoryID int        `json:"category_id"`
	Title      string     `json:"title"`
	SeoURL     string     `json:"seo_url"`
	Body       string     `json:"body"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	Tags       []Tag      `json:"tags"`
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
