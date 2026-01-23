package dynamodbx

import (
	"context"
	"errors"
	"strconv"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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
		ConditionExpression: aws.String("attribute_not_exists(SessionID)"),
	}

	_, putErr := DB.PutItem(ctx, putParams)
	if putErr != nil {
		return putErr
	}

	return nil
}

func GetChatSessionBySessionID(ctx context.Context, sessionID string) (chatSession models.ChatSession, error error) {
	getParams := &dynamodb.GetItemInput{
		TableName: aws.String("ChatSession"),
		Key: map[string]types.AttributeValue{
			"SessionID": &types.AttributeValueMemberS{Value: sessionID},
		},
	}

	resp, getErr := DB.GetItem(ctx, getParams)
	if getErr != nil {
		return models.ChatSession{}, getErr
	}
	if resp.Item == nil {
		return models.ChatSession{}, errors.New("Session doesnt exist")
	}

	var session models.ChatSession

	unmarshallErr := attributevalue.UnmarshalMap(resp.Item, &session)
	if unmarshallErr != nil {
		return models.ChatSession{}, unmarshallErr
	}

	return session, nil
}
