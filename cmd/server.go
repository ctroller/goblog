package main

import (
	"goblog/internal/config"
	"goblog/internal/handler"
	"goblog/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	config := config.BlogConfig{
		PostService: service.NewPostService(),
	}

	RegisterRoutes(router, config)

	println("Server is running on http://localhost:8080")
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func RegisterRoutes(router *mux.Router, config config.BlogConfig) {
	// Serve static files
	fs := http.FileServer(http.Dir("ui/static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	router.PathPrefix("/").HandlerFunc(handler.RootHandler(config))
}
