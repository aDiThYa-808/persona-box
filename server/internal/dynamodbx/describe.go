package dynamodbx

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DescribeTable(ctx context.Context, name string) error {
	_, err := DB.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(name),
	})

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
