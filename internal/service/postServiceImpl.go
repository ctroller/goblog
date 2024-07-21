package service

import (
	"errors"
	"goblog/internal/dto"
	"goblog/internal/rest"
	"net/http"
	"time"
)

type postServiceImpl struct{}

func NewPostService() PostService {
	return &postServiceImpl{}
}

var posts = []dto.Post{
	{
		ID:        1,
		Title:     "Hello, World!",
		SeoURL:    "hello-world",
		Body:      "This is my first post.",
		CreatedAt: time.Now(),
	},
}

func (s *postServiceImpl) GetBySeoURL(seoURL string) (*dto.Post, error) {
	for _, p := range posts {
		if p.SeoURL == seoURL {
			return &p, nil
		}
	}
	return nil, nil
}

func (s *postServiceImpl) GetByFilter(filter Filter, page *int, perPage *int) (Posts, error) {
	return Posts{
		Posts:      posts,
		Pagination: rest.Pagination{},
	}, nil
}

func (s *postServiceImpl) GetCategories(max *int) ([]dto.Category, error) {
	return nil, errors.New("not implemented")
}

func (s *postServiceImpl) GetTags(max *int) ([]dto.Tag, error) {
	return nil, errors.New("not implemented")
}

func (s *postServiceImpl) GetArchives(max *int) ([]Archive, error) {
	return nil, errors.New("not implemented")
}

func (s *postServiceImpl) GetFilterParams(request *http.Request) Filter {
	return Filter{}
}
