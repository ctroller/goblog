package service

import (
	"errors"
	"goblog/internal/block"
	"goblog/internal/dto"
	"time"
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
		Blocks: []block.ContentBlock{
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

func (s *postServiceImpl) GetByFilter(filter PostFilter, page *int, perPage *int) ([]dto.Post, error) {
	return posts, nil
}

func (s *postServiceImpl) GetCategories(max *int) ([]dto.Category, error) {
	return nil, errors.New("not implemented")
}

func (s *postServiceImpl) GetDates(max *int) ([]PostDate, error) {
	return nil, errors.New("not implemented")
}

func (s *postServiceImpl) GetAll() ([]dto.Post, error) {
	return posts, nil
}
