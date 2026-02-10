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

/*
Creates a new item in the ChatSession table if it doesnt already exist.
Returns error if it fails to create the item.
*/
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

/*
Returns the item from ChatSession table that has the same SessionID and PersonaID as provided.
*/
func GetChatSessionByID(ctx context.Context, personaID string, sessionID string) (chatSession models.ChatSession, error error) {
	getParams := &dynamodb.GetItemInput{
		TableName: aws.String("ChatSession"),
		Key: map[string]types.AttributeValue{
			"SessionID": &types.AttributeValueMemberS{Value: sessionID},
			"PersonaID": &types.AttributeValueMemberS{Value: personaID},
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

func GetAllChatSessionsOfPersona(ctx context.Context, personaid string) ([]models.ChatSessionList, error) {
	queryParams := &dynamodb.QueryInput{
		TableName:              aws.String("ChatSession"),
		KeyConditionExpression: aws.String("PersonaID = :pid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pid": &types.AttributeValueMemberS{Value: personaid},
		},
	}
	resp, queryErr := DB.Query(ctx, queryParams)
	if queryErr != nil {
		return []models.ChatSessionList{}, queryErr
	}
	if resp.Items == nil {
		return []models.ChatSessionList{}, errors.New("no chat sessions found")
	}

	chatSessions := make([]models.ChatSessionList, len(resp.Items))

	for i, item := range resp.Items {
		unmarshallErr := attributevalue.UnmarshalMap(item, &chatSessions[i])
		if unmarshallErr != nil {
			return []models.ChatSessionList{}, unmarshallErr
		}
	}

	return chatSessions, nil
}

/*
Updates 'UpdatedAt', 'MessageCount' and 'TokenCount' of the ChatSession item with the provided SessionID.
Returns an error if update fails.
*/
func UpdateSessionMessageAndTokenCount(ctx context.Context, personaID string, sessionID string, updatedAt string, messageCount int, tokensUsed int) error {
	updateExpression := "SET UpdatedAt = :now ADD MessageCount :msgcount, TokenCount :tkncount"
	attributeValues := map[string]types.AttributeValue{
		":now":      &types.AttributeValueMemberS{Value: updatedAt},
		":msgcount": &types.AttributeValueMemberN{Value: strconv.Itoa(messageCount)},
		":tkncount": &types.AttributeValueMemberN{Value: strconv.Itoa(tokensUsed)},
	}

	updateParams := &dynamodb.UpdateItemInput{
		TableName: aws.String("ChatSession"),
		Key: map[string]types.AttributeValue{
			"SessionID": &types.AttributeValueMemberS{Value: sessionID},
			"PersonaID": &types.AttributeValueMemberS{Value: personaID},
		},
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: attributeValues,
		ConditionExpression:       aws.String("attribute_exists(SessionID)"),
	}

	_, updateErr := DB.UpdateItem(ctx, updateParams)
	if updateErr != nil {
		return updateErr
	}

	return nil
}

func DeleteSession(ctx context.Context, personaid string, sessionid string) error {
	deleteMsgErr := DeleteAllMessagesOfASession(ctx, sessionid)
	if deleteMsgErr != nil {
		return deleteMsgErr
	}

	deleteParams := &dynamodb.DeleteItemInput{
		TableName: aws.String("ChatSession"),
		Key: map[string]types.AttributeValue{
			"PersonaID": &types.AttributeValueMemberS{Value: personaid},
			"SessionID": &types.AttributeValueMemberS{Value: sessionid},
		},
	}

	_, deleteErr := DB.DeleteItem(ctx, deleteParams)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func DeleteAllSessionsOfAPersona(ctx context.Context, personaid string) error {
	queryParams := &dynamodb.QueryInput{
		TableName:              aws.String("ChatSession"),
		KeyConditionExpression: aws.String("PersonaID = :pid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pid": &types.AttributeValueMemberS{Value: personaid},
		},
		ProjectionExpression: aws.String("PersonaID, SessionID"),
	}

	resp, queryErr := DB.Query(ctx, queryParams)
	if queryErr != nil {
		return queryErr
	}

	for _, item := range resp.Items {

		//first delete all the messages of the session
		var sessionid string
		unmarshalErr := attributevalue.Unmarshal(item["SessionID"], &sessionid)
		if unmarshalErr != nil {
			return unmarshalErr
		}
		deleteMsgErr := DeleteAllMessagesOfASession(ctx, sessionid)
		if deleteMsgErr != nil {
			return deleteMsgErr
		}

		// then delete the session itself
		deleteParams := &dynamodb.DeleteItemInput{
			TableName: aws.String("ChatSession"),
			Key: map[string]types.AttributeValue{
				"PersonaID": item["PersonaID"],
				"SessionID": item["SessionID"],
			},
		}
		_, deleteErr := DB.DeleteItem(ctx, deleteParams)
		if deleteErr != nil {
			return deleteErr
		}
	}

	return nil
}
