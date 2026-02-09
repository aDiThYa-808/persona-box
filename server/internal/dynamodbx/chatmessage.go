package dynamodbx

import (
	"context"
	"errors"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func StoreChatMessage(ctx context.Context, message models.ChatMessage) error {
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("ChatMessage"),
		Item: map[string]types.AttributeValue{
			"SessionID": &types.AttributeValueMemberS{Value: message.SessionID},
			"CreatedAt": &types.AttributeValueMemberS{Value: message.CreatedAt},
			"Role":      &types.AttributeValueMemberS{Value: message.Role},
			"Message":   &types.AttributeValueMemberS{Value: message.Message},
		},
	}

	_, putErr := DB.PutItem(ctx, putParams)

	if putErr != nil {
		return putErr
	}

	return nil
}

func GetAllSessionMessages(ctx context.Context, sessionid string) ([]models.ChatMessage, error) {
	queryParams := &dynamodb.QueryInput{
		TableName:              aws.String("ChatMessage"),
		KeyConditionExpression: aws.String("SessionID = :sid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":sid": &types.AttributeValueMemberS{Value: sessionid},
		},
	}

	resp, queryErr := DB.Query(ctx, queryParams)
	if queryErr != nil {
		return []models.ChatMessage{}, queryErr
	}
	if resp.Items == nil {
		return []models.ChatMessage{}, errors.New("no messages found")
	}

	messages := make([]models.ChatMessage, len(resp.Items))

	for i, item := range resp.Items {
		unmarshalErr := attributevalue.UnmarshalMap(item, &messages[i])
		if unmarshalErr != nil {
			return []models.ChatMessage{}, unmarshalErr
		}
	}

	return messages, nil
}

func DeleteAllMessagesOfASession(ctx context.Context, sessionid string) error {

	queryParams := &dynamodb.QueryInput{
		TableName:              aws.String("ChatMessage"),
		KeyConditionExpression: aws.String("SessionID = :sid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":sid": &types.AttributeValueMemberS{Value: sessionid},
		},
		ProjectionExpression: aws.String("SessionID, CreatedAt"),
	}

	resp, queryErr := DB.Query(ctx, queryParams)
	if queryErr != nil {
		return queryErr
	}

	for _, item := range resp.Items {
		deleteParams := &dynamodb.DeleteItemInput{
			TableName: aws.String("ChatMessage"),
			Key: map[string]types.AttributeValue{
				"SessionID": item["SessionID"],
				"CreatedAt": item["CreatedAt"],
			},
		}
		_, deleteErr := DB.DeleteItem(ctx, deleteParams)
		if deleteErr != nil {
			return deleteErr
		}
	}

	return nil
}
