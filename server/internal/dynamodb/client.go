package dynamodb

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DB *dynamodb.Client

func Init() {
	ctx := context.Background()

	cfg, cfgErr := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-north-1"),
	)

	if cfgErr != nil {
		log.Println("Failed to create dynamodb client")
		panic(cfgErr)
	}

	DB = dynamodb.NewFromConfig(cfg)
}
