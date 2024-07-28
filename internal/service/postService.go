package service

import (
	"goblog/internal/dto"
	"net/http"
)

// Date is a Year / Month representation of Posts
type Date struct {
	Year  *int
	Month *int
}

// Filter defines criteria for filtering posts. It supports filtering by archive date, tags, category, number of posts, and posts since a specific date (used for pagination).
type Filter struct {
	Date       *Date
	Tags       *[]string
	CategoryID *int
	Num        *int
}

// PostService defines the interface for services that manage blog posts. It includes methods for retrieving posts by SEO URL, filtering posts, and getting categories, tags, and archives.
type PostService interface {
	GetAll() ([]dto.Post, error)

	// GetBySeoURL retrieves a single post by its SEO-friendly URL.
	GetBySeoURL(seoURL string) (*dto.Post, error)

	// GetByFilter retrieves posts based on various filter criteria, including pagination parameters.
	GetByFilter(filter Filter, page *int, perPage *int) ([]dto.Post, error)

	// GetCategories retrieves categories used in the blog
	GetCategories(max *int) ([]dto.Category, error)

	// GetDates retrieves a list of used dates (grouped by year and month)
	GetDates(max *int) ([]Date, error)

	// GetFilterParams extracts filter criteria from an HTTP request
	GetFilterParams(request *http.Request) *Filter
}
