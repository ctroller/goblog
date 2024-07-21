package service

import (
	"goblog/internal/dto"
	"goblog/internal/rest"
	"net/http"
	"time"
)

// Archive is a Year / Month representation of Posts
type Archive struct {
	Year  *int
	Month *int
}

// Filter defines criteria for filtering posts. It supports filtering by archive date, tags, category, number of posts, and posts since a specific date (used for pagination).
type Filter struct {
	Archive    *Archive
	Tags       *[]string
	CategoryID *int
	Num        *int
	Since      *time.Time
}

// Pagination wrapper for a list of posts
type Posts struct {
	Posts      []dto.Post      `json:"posts"`
	Pagination rest.Pagination `json:"pagination"`
}

// PostService defines the interface for services that manage blog posts. It includes methods for retrieving posts by SEO URL, filtering posts, and getting categories, tags, and archives.
type PostService interface {
	// GetBySeoURL retrieves a single post by its SEO-friendly URL.
	GetBySeoURL(seoURL string) (*dto.Post, error)

	// GetByFilter retrieves posts based on various filter criteria, including pagination parameters.
	GetByFilter(filter Filter, page *int, perPage *int) (Posts, error)

	// GetCategories retrieves categories used in the blog
	GetCategories(max *int) ([]dto.Category, error)

	// GetTags retrieves tags used in the blog
	GetTags(max *int) ([]dto.Tag, error)

	// GetArchives retrieves a list of archives (grouped by year and month)
	GetArchives(max *int) ([]Archive, error)

	// GetFilterParams extracts filter criteria from an HTTP request
	GetFilterParams(request *http.Request) Filter
}
