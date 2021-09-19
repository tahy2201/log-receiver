package main

import (
	"net/http"

	"log-receiver/internal/controller"

	"github.com/rs/cors"
	"google.golang.org/appengine"
)

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/traceloga", controller.TracelogA)
	handler := cors.Default().Handler(mux)

	// Handle all requests using net/http
	http.Handle("/", handler)
}

func main() {
	http.ListenAndServe(":8080", nil)
	appengine.Main()
}
