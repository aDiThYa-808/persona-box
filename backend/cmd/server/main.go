package main

import (
	"log"
	"net/http"

	handlers "github.com/aDiThYa-808/persona-box/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))

	log.Println("Server Running in port 3000...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
