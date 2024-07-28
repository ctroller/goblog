package service

import (
	"goblog/internal/dto"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/rs/zerolog/log"
)

// PostDate is a Year / Month representation of Posts
type PostDate struct {
	Year  *int
	Month *int
}

// PostFilter defines criteria for filtering posts. It supports filtering by archive date, tags, category, number of posts, and posts since a specific date (used for pagination).
type PostFilter struct {
	Date       *PostDate
	Tags       *[]string
	CategoryID *int
	Num        *int
}

// PostService defines the interface for services that manage blog posts. It includes methods for retrieving posts by SEO URL, filtering posts, and getting categories, tags, and archives.
type PostService interface {
	// Returns all posts
	GetAll() ([]dto.Post, error)

	// GetByFilter retrieves posts based on various filter criteria, including pagination parameters.
	GetByFilter(filter PostFilter, page *int, perPage *int) ([]dto.Post, error)

	// GetCategories retrieves categories used in the blog
	GetCategories(max *int) ([]dto.Category, error)

	// GetDates retrieves a list of used dates (grouped by year and month)
	GetDates(max *int) ([]PostDate, error)
}

func GetFilterParams(request *http.Request) *PostFilter {
	var decoder = schema.NewDecoder()
	var filter *PostFilter
	err := decoder.Decode(&filter, request.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode filter params")
		return nil
	}

	return filter
}
