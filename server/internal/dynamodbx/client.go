package dynamodbx

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Global dynamodb client variable
var DB *dynamodb.Client

// Initializes dynamodb client
func InitializeDynamoDBClient() {
	ctx := context.Background()

	cfg, cfgErr := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-north-1"),
		config.WithRetryMaxAttempts(3),
	)

	if cfgErr != nil {
		log.Println("Failed to create dynamodb client")
		panic(cfgErr)
	}

	DB = dynamodb.NewFromConfig(cfg)
}
