package cache

import (
	"errors"
	"goblog/internal/block"
	"goblog/internal/config"
	"goblog/internal/dto"
	"goblog/internal/service"
	"os"
	"testing"
	"time"
)

// Mock the necessary dependencies
type MockPostService struct{}

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

func (m *MockPostService) GetAll() ([]dto.Post, error) {
	return posts, nil
}

func (s *MockPostService) GetByFilter(filter service.PostFilter, page *int, perPage *int) ([]dto.Post, error) {
	return posts, nil
}

func (s *MockPostService) GetCategories(max *int) ([]dto.Category, error) {
	return nil, errors.New("not implemented")
}

func (s *MockPostService) GetDates(max *int) ([]service.PostDate, error) {
	return nil, errors.New("not implemented")
}

func initCacheDir(b config.BlogConfig) {
	tmpdir := os.TempDir()
	b.PostCacheConfig.CacheDir = tmpdir + "/goblog-test-cache"
}

func TestCacheAllPosts(t *testing.T) {
	mockConfig := config.BlogConfig{
		PostService: &MockPostService{},
	}
	initCacheDir(mockConfig)

	CacheAllPosts(mockConfig)

	// Check if the cache files were created
	for _, post := range posts {
		cacheFile := getPostCacheFilePath(&post, mockConfig.PostCacheConfig)
		if _, err := os.Stat(cacheFile); err != nil {
			t.Errorf("Cache file %s does not exist", cacheFile)
		}
	}

	// cleanup
	for _, post := range posts {
		cacheFile := getPostCacheFilePath(&post, mockConfig.PostCacheConfig)
		os.Remove(cacheFile)
	}
}

func TestCachePost(t *testing.T) {
	mockConfig := config.BlogConfig{}
	initCacheDir(mockConfig)
	
	post := posts[0]

	CachePost(&post, mockConfig.PostCacheConfig)

	cacheFile := getPostCacheFilePath(&post, mockConfig.PostCacheConfig)
	if _, err := os.Stat(cacheFile); err != nil {
		t.Errorf("Cache file %s does not exist", cacheFile)
	}

	// cleanup
	os.Remove(cacheFile)
}
