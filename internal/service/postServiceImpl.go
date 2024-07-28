package service

import (
	"errors"
	"goblog/internal/block"
	"goblog/internal/dto"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/rs/zerolog/log"
)

type postServiceImpl struct{}

func NewPostService() PostService {
	return &postServiceImpl{}
}

var posts = []dto.Post{
	{
		ID:      1,
		Title:   "Hello, World!",
		Summary: "A simple introduction to the world",
		SeoURL:  "hello-world",
		Body: []block.ContentBlock{
			block.NewTextBlock("Hello, World!"),
			block.NewCodeBlock(`package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`, "go"),
			block.NewCodeBlock(`import java.util.*;

public class Main {
		public static void main(String... args) {
			System.out.println("Hello, World!                                      LOOOOOOOOOOOOOOOOOOOONG");
		}
}`, "java"),
		},
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

func (s *postServiceImpl) GetByFilter(filter Filter, page *int, perPage *int) ([]dto.Post, error) {
	return posts, nil
}

func (s *postServiceImpl) GetCategories(max *int) ([]dto.Category, error) {
	return nil, errors.New("not implemented")
}

func (s *postServiceImpl) GetDates(max *int) ([]Date, error) {
	return nil, errors.New("not implemented")
}

func (s *postServiceImpl) GetFilterParams(request *http.Request) *Filter {
	var decoder = schema.NewDecoder()
	var filter *Filter
	err := decoder.Decode(&filter, request.URL.Query())
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode filter params")
		return nil
	}

	return filter
}

func (s *postServiceImpl) GetAll() ([]dto.Post, error) {
	return posts, nil
}
