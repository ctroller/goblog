package main

import (
	"flag"
	"goblog/internal/config"
	"goblog/internal/handler"
	"goblog/internal/service"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	router := mux.NewRouter()

	config := config.BlogConfig{
		PostService: service.NewPostService(),
	}
	Configure()

	RegisterRoutes(router, config)

	println("Server is running on http://localhost:8080")
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal().Err(srv.ListenAndServe()).Msg("")
}

func RegisterRoutes(router *mux.Router, config config.BlogConfig) {
	// Serve static files
	fs := http.FileServer(http.Dir("ui/static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.NotFoundHandler = handler.NotFoundHandler(config)
	router.PathPrefix("/posts/").HandlerFunc(handler.PostDetailHandler(config))
	router.PathPrefix("/").HandlerFunc(handler.RootHandler(config))
}

var logLevelLookup = map[string]zerolog.Level{
	"debug":    zerolog.DebugLevel,
	"info":     zerolog.InfoLevel,
	"warn":     zerolog.WarnLevel,
	"error":    zerolog.ErrorLevel,
	"fatal":    zerolog.FatalLevel,
	"panic":    zerolog.PanicLevel,
	"disabled": zerolog.Disabled,
}

func Configure() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logLevel := flag.String("loglevel", "warn", "sets the global log level")
	flag.Parse()
	level, ok := logLevelLookup[*logLevel]
	if !ok {
		level = zerolog.WarnLevel
	}

	zerolog.SetGlobalLevel(level)
}
