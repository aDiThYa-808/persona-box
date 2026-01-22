package dynamodbx

import (
	"context"
	"strconv"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateChatMessage(ctx context.Context, message models.ChatMessage) error {
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("ChatMessage"),
		Item: map[string]types.AttributeValue{
			"SessionID":    &types.AttributeValueMemberS{Value: message.SessionID},
			"CreatedAt":    &types.AttributeValueMemberS{Value: message.CreatedAt},
			"Sender":       &types.AttributeValueMemberS{Value: message.Sender},
			"Content":      &types.AttributeValueMemberS{Value: message.Content},
			"TokensUsed":   &types.AttributeValueMemberN{Value: strconv.Itoa(message.TokensUsed)},
			"ModelVersion": &types.AttributeValueMemberS{Value: message.ModelVersion},
		},
	}

	_, putErr := DB.PutItem(ctx, putParams)

	if putErr != nil {
		return putErr
	}

	return nil
}
