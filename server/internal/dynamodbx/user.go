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

/*
Creates a new item in the User table using PutItem().
Returns created = true if item is added successfully.
Returns created = false if item already exists with the same user.UserID.
Returns error is PutItem() fails.
*/
func CreateNewUser(ctx context.Context, user models.User) (created bool, error error) {

	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("User"),
		Item: map[string]types.AttributeValue{
			"UserID":      &types.AttributeValueMemberS{Value: user.UserID},
			"Email":       &types.AttributeValueMemberS{Value: user.Email},
			"DisplayName": &types.AttributeValueMemberS{Value: user.DisplayName},
			"CreatedAt":   &types.AttributeValueMemberS{Value: user.CreatedAt},
		},
		ConditionExpression: aws.String("attribute_not_exists(UserID)"),
	}

	_, putErr := DB.PutItem(ctx, putParams)

	if putErr != nil {
		//error type for conditional expression not satisfied
		var ccfe *types.ConditionalCheckFailedException

		if errors.As(putErr, &ccfe) {
			//creation fails but no error. expected behaviour
			return false, nil
		}
		return false, putErr
	}

	return true, nil
}

/*
Gets the item from User table that has the same partition key value as the provided userID using GetItem().
Return user if it is found.
Returns error if the item doesnt exist or GetItem() fails.
*/
func GetUser(ctx context.Context, userID string) (user models.User, error error) {
	getParams := &dynamodb.GetItemInput{
		TableName: aws.String("User"),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: userID},
		},
	}

	resp, getErr := DB.GetItem(ctx, getParams)

	if getErr != nil {
		return models.User{}, getErr
	}
	if resp.Item == nil {
		return models.User{}, errors.New("user doesnt exist")
	}

	var usr models.User

	//reads the resp.Item map and fills the feilds of usr.
	unmarshalErr := attributevalue.UnmarshalMap(resp.Item, &usr)
	if unmarshalErr != nil {
		return models.User{}, unmarshalErr
	}
	return usr, nil
}
