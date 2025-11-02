package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/joho/godotenv"

	handlers "github.com/aDiThYa-808/persona-box/internal/handlers"
	middlewares "github.com/aDiThYa-808/persona-box/internal/middlewares"
)

// helpers
func isLambda() bool {
	// Check for an aws env which will not be present locally.
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("env not found.")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))
	mux.Handle("/chat", http.HandlerFunc(handlers.ChatHandler))

	// add cors headers
	handler := middlewares.StripStagePrefix(mux)
	handler = middlewares.WithCors(handler)

	// local environment
	if !isLambda() {
		log.Println("Server Running locally in port 3000...")
		log.Fatal(http.ListenAndServe(":3000", handler))
	} else {
		adapter := httpadapter.NewV2(handler)
		lambda.Start(adapter.ProxyWithContext)
	}

}
