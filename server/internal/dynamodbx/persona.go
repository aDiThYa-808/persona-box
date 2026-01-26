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
Creates a new item in the Persona table using PutItem().
Returns error if PutItem() fails.
*/
func CreateNewPersona(ctx context.Context, persona models.Persona) error {
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("Persona"),
		Item: map[string]types.AttributeValue{
			"UserID":             &types.AttributeValueMemberS{Value: persona.UserID},
			"PersonaID":          &types.AttributeValueMemberS{Value: persona.PersonaID},
			"PersonaName":        &types.AttributeValueMemberS{Value: persona.PersonaName},
			"PersonaDescription": &types.AttributeValueMemberS{Value: persona.PersonaDescription},
			"Age":                &types.AttributeValueMemberN{Value: strconv.Itoa(persona.Age)},
			"Pronouns":           &types.AttributeValueMemberSS{Value: persona.Pronouns},
			"Openness":           &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.Openness), 'f', -1, 32)},
			"Conscientiousness":  &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.Conscientiousness), 'f', -1, 32)},
			"Extraversion":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.Extraversion), 'f', -1, 32)},
			"Agreeableness":      &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.Agreeableness), 'f', -1, 32)},
			"Neuroticism":        &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.Neuroticism), 'f', -1, 32)},
			"Intelligence":       &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.Intelligence), 'f', -1, 32)},
			"ThinkingStyle":      &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.ThinkingStyle), 'f', -1, 32)},
			"Tone":               &types.AttributeValueMemberSS{Value: persona.Tone},
			"HumorLevel":         &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.HumorLevel), 'f', -1, 32)},
			"MoodFluctuation":    &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.MoodFluctuation), 'f', -1, 32)},
			"Likes":              &types.AttributeValueMemberSS{Value: persona.Likes},
			"Dislikes":           &types.AttributeValueMemberSS{Value: persona.Dislikes},
			"Formality":          &types.AttributeValueMemberN{Value: strconv.FormatFloat(float64(persona.Formality), 'f', -1, 32)},
			"Fluency":            &types.AttributeValueMemberS{Value: persona.Fluency},
			"EmojiUsage":         &types.AttributeValueMemberS{Value: persona.EmojiUsage},
			"ResponseLength":     &types.AttributeValueMemberS{Value: persona.ResponseLength},
			"CreatedAt":          &types.AttributeValueMemberS{Value: persona.CreatedAt},
		},
		ConditionExpression: aws.String("attribute_not_exists(PersonaID)"),
	}

	_, putErr := DB.PutItem(ctx, putParams)
	if putErr != nil {
		return putErr
	}

	return nil
}

/*
Returns all the items from Persona table that has the same partition key value as the provided userID using Query().
Returns error if query fails.
*/
func GetUsersPersonas(ctx context.Context, userID string) ([]models.PersonaList, error) {
	queryParams := &dynamodb.QueryInput{
		TableName:              aws.String("Persona"),
		KeyConditionExpression: aws.String("UserID = :uid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":uid": &types.AttributeValueMemberS{Value: userID},
		},
	}

	resp, queryErr := DB.Query(ctx, queryParams)
	if queryErr != nil {
		return []models.PersonaList{}, queryErr
	}

	personas := make([]models.PersonaList, len(resp.Items))

	for i, item := range resp.Items {
		unmarshallErr := attributevalue.UnmarshalMap(item, &personas[i])
		if unmarshallErr != nil {
			return []models.PersonaList{}, unmarshallErr
		}
	}

	return personas, nil
}

/*
Returns the item from Persona table that has the same PersonID and UserID primary key as the ones provided.
*/
func GetPersonaByPersonaID(ctx context.Context, userID string, personaID string) (models.Persona, error) {
	getParams := &dynamodb.GetItemInput{
		TableName: aws.String("Persona"),
		Key: map[string]types.AttributeValue{
			"PersonaID": &types.AttributeValueMemberS{Value: personaID},
			"UserID":    &types.AttributeValueMemberS{Value: userID},
		},
	}

	resp, getErr := DB.GetItem(ctx, getParams)
	if getErr != nil {
		return models.Persona{}, getErr
	}
	if resp.Item == nil {
		return models.Persona{}, errors.New("persona doesnt exist")
	}

	persona := models.Persona{}
	unmarshallErr := attributevalue.UnmarshalMap(resp.Item, &persona)
	if unmarshallErr != nil {
		return models.Persona{}, unmarshallErr
	}

	return persona, nil
}
