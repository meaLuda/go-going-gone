package user

  import (
    "encoding/json"
    "errors"

    "aws-lambda-in-go-lang/pkg/validators"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// ERROR LIST
var (
    ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
    ErrorFailedToFetchRecord     = "failed to fetch record"
    ErrorInvalidUserData         = "invalid user data"
    ErrorInvalidEmail            = "invalid email"
    ErrorCouldNotMarshalItem     = "could not marshal item"
    ErrorCouldNotDeleteItem      = "could not delete item"
    ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
    ErrorUserAlreadyExists       = "user.User already exists"
    ErrorUserDoesNotExists       = "user.User does not exist"
)

type User struct{
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
            "email": {
                S: aws.String(email),
            },
        },
        TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(input)
    if err != nil {
        return nil, errors.New(ErrorFailedToFetchRecord)

    }

    item := new(User)
    err = dynamodbattribute.UnmarshalMap(result.Item, item)
    if err != nil {
        return nil, errors.New(ErrorFailedToUnmarshalRecord)
    }
    return item, nil
}

