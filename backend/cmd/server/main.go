package main

import (
	"log"
	"net/http"

	handlers "github.com/aDiThYa-808/persona-box/internal/handlers"
	middlewares "github.com/aDiThYa-808/persona-box/internal/middlewares"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))

	handler := middlewares.WithCors(mux)

	log.Println("Server Running in port 3000...")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
