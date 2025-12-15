package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MicahParks/keyfunc"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/joho/godotenv"

	"github.com/aDiThYa-808/persona-box/internal/dynamodb"
	handlers "github.com/aDiThYa-808/persona-box/internal/handlers"
	middlewares "github.com/aDiThYa-808/persona-box/internal/middlewares"
)

// helpers
func isLambda() bool {
	// Check for an aws env which will not be present locally.
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}

var (
	googleJWKS *keyfunc.JWKS
)

func init() {

	//load ENV in local environment
	if !isLambda() {
		if loadErr := godotenv.Load(); loadErr != nil {
			log.Fatal("env not found." + loadErr.Error())
		}
	}

	//create dynamodb client
	dynamodb.Init()

	// Fetch and hold google's JWKS for verification
	var jwksErr error
	googleJWKS, jwksErr = keyfunc.Get("https://www.googleapis.com/oauth2/v3/certs", keyfunc.Options{})
	if jwksErr != nil {
		log.Fatal("Failed to fetch JWKS: " + jwksErr.Error())
	}
}

func main() {
	mux := http.NewServeMux()

	h := handlers.New(googleJWKS)

	mux.Handle("/auth/google", http.HandlerFunc(h.GoogleAuthHandler))
	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))
	mux.Handle("/chat", http.HandlerFunc(handlers.ChatHandler))
	mux.Handle("/test", http.HandlerFunc(handlers.Tstendpoint))

	//remove prefix from incoming req url
	handler := middlewares.StripStagePrefix(mux)

	// add cors headers
	handler = middlewares.WithCors(handler)

	// local or Prod environment check
	if !isLambda() {
		log.Println("Server Running locally in port 3000...")
		log.Fatal(http.ListenAndServe(":3000", handler))
	} else {
		adapter := httpadapter.NewV2(handler)
		lambda.Start(adapter.ProxyWithContext)
	}

}
