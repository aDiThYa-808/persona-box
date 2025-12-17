package dynamodbx

import (
	"context"
	"errors"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

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
