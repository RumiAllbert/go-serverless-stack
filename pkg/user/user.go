package user

import (
	"enconding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/dynamodb"
	"github.com/aws/aws-sdk-go/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var(
	ErrorFailedToFetchRecord = "failed to fetch record"
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
	
}

func Createuser() {

}

func UpdateUser() {

}

func Deleteuser() error {

}
