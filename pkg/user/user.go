package user

import (
	"github.com/rumiallbert/go-serverless-stack/pkg/user"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var(
	ErrorFailedToFetchRecord = "failed to fetch record"
	ErrorInvalidUserData = "invalid user data"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrroInvalidEmail = "invalid email"
	ErrorCouldNotMarshalItem = "could not marshal item"
	ErrorCouldNotDeleteItem = "could not delete item"
	ErrorCouldNotDynamoPutItem = "could not dynamo put item"
	ErrorUserAlreadyExists = "user.User already exists"
	ErrorUserDoesNotExist = "user.User does not exist"
)

type User struct {
	Email		string `json:"email"`
	FirstName	string `json:"firstName"`
	LastName	string `json:"lastName"`
}

func FetchUser(email, tableName string, dynaClient dynomodbiface.DynamoDBAPI)(*User, error) {
	
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynomodb.AttributeValue{
			"email": {
				S: aws.String(email)
			}
		},
		TableName: aws.String(tableName)
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new(user)
	err = dynomodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*[]Users) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName)
	}
	results, err:= dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]User)
	err = dynomodbattribute.UnmarshalMap(result.Items, item)
	return item, nil
}

func Createuser(req events.APIGatewayProxyRequest, tablName string, dynaClient dynomodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error)
(*User, error){
	var u User

	if err :=json.Unmarshal([]byte(req.Body), &u): if err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}
	// First check whether the email is valid
	if !validtors.IsEmailValid(u.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}
	// Check whether the user already exists
	currentUser, _ := FetchUser(u.Email, tableName, dynaClient)
	if currenUser != nil && len(CurrentUser.email) != 0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	}
	// If user doesn't exist and email is valid, save the user
	av, erro :=	dynamodbattribute.marshalMap(u)
	if err != nil {
		return nil, errors.new(ErrorCouldNotMarshalItem)
	}

	input: &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(tableName)
	}

	_, err =dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &u, nil
}

}

func UpdateUser() {

}

func Deleteuser() error {

}
