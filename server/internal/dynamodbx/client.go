package dynamodbx

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Global dynamodb client variable
var DB *dynamodb.Client

// Initializes dynamodb client
func InitializeDynamoDBClient() {
	ctx := context.Background()

	region := os.Getenv("AWS_REGION")
	if region == "" {
		panic("aws region not found")
	}

	cfg, cfgErr := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithRetryMaxAttempts(3),
	)

	if cfgErr != nil {
		log.Println("Failed to create dynamodb client")
		panic(cfgErr)
	}

	DB = dynamodb.NewFromConfig(cfg)
}
