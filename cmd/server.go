package main

import (
	"flag"
	"goblog/internal/cache"
	"goblog/internal/config"
	"goblog/internal/handler"
	"goblog/internal/service"
	"net/http"
	"os"
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

	configure()
	cache.CacheAllPosts(config)
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
	router.NotFoundHandler = handler.NotFoundHandler(config)

	// Serve static files
	fs := http.FileServer(http.Dir("ui/static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Post cache
	router.PathPrefix("/posts").Handler(http.StripPrefix("/posts/", Static404Handler(http.Dir("cache/posts"), router)))

	//router.PathPrefix("/posts/").HandlerFunc(handler.PostDetailHandler(config))
	router.Path("/").HandlerFunc(handler.RootHandler(config))
}

func Static404Handler(root http.FileSystem, router *mux.Router) http.Handler {
	fs := http.FileServer(root)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := root.Open(r.URL.Path)
		if err != nil && os.IsNotExist(err) {
			router.NotFoundHandler.ServeHTTP(w, r)
		} else {
			if err == nil {
				f.Close()
			}
			fs.ServeHTTP(w, r)
		}
	})
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

func configure() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logLevel := flag.String("loglevel", "warn", "sets the global log level")
	flag.Parse()
	level, ok := logLevelLookup[*logLevel]
	if !ok {
		level = zerolog.WarnLevel
	}

	zerolog.SetGlobalLevel(level)
}
