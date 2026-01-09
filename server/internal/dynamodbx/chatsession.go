package dynamodbx

import (
	"context"
	"strconv"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateNewChatSession(ctx context.Context, session models.ChatSession) error {
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("ChatSession"),
		Item: map[string]types.AttributeValue{
			"SessionID":    &types.AttributeValueMemberS{Value: session.SessionID},
			"PersonaID":    &types.AttributeValueMemberS{Value: session.PersonaID},
			"Title":        &types.AttributeValueMemberS{Value: session.Title},
			"CreatedAt":    &types.AttributeValueMemberS{Value: session.CreatedAt},
			"UpdatedAt":    &types.AttributeValueMemberS{Value: session.UpdatedAt},
			"MessageCount": &types.AttributeValueMemberN{Value: strconv.Itoa(session.MessageCount)},
			"TokenCount":   &types.AttributeValueMemberN{Value: strconv.Itoa(session.TokenCount)},
			"Summary":      &types.AttributeValueMemberS{Value: session.Summary},
		},
	}

	_, putErr := DB.PutItem(ctx, putParams)
	if putErr != nil {
		return putErr
	}

	return nil
}
