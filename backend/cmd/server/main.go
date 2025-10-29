package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	handlers "github.com/aDiThYa-808/persona-box/internal/handlers"
	middlewares "github.com/aDiThYa-808/persona-box/internal/middlewares"
)

// helpers
func isLambda() bool {
	// Check for an aws env which will not be present locally.
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))

	// add cors headers
	handler := middlewares.WithCors(mux)

	// local environment
	if !isLambda() {
		log.Println("Server Running locally in port 3000...")
		log.Fatal(http.ListenAndServe(":3000", handler))
	}

	adapter := httpadapter.New(handler)
	lambda.Start(adapter.ProxyWithContext)

}
