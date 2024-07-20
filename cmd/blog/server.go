package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve static files
	fs := http.FileServer(http.Dir("ui/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
