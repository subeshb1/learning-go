package model

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

// Person Model
type Person struct {
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
	Name   string `json:"name"`
}

func (p Person) Update() (Person, error) {

	return Person{Age: 10, Gender: 1, Name: "Subesh"}, nil
}

// People Model
type People struct {
}

func (p People) Where(svc *dynamodb.DynamoDB) (*dynamodb.QueryOutput, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String("Subesh"),
			},
		},
		KeyConditionExpression: aws.String("Name = :v1"),
		ProjectionExpression:   aws.String("SongTitle"),
		TableName:              aws.String("Person"),
	}
	result, _ := svc.Query(input)
	return result, nil
}

func (p People) Find() (Person, error) {

	return Person{Age: 10, Gender: 1, Name: "Subesh"}, nil
}

func (p People) Create(svc *dynamodb.DynamoDB, name string, age, gender int) (*dynamodb.PutItemOutput, error) {
	uid := fmt.Sprintf("%s", uuid.NewV4())
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"UID": {
				S: aws.String(uid),
			},
			"Name": {
				S: aws.String(name),
			},
			"Age": {
				N: aws.String(strconv.Itoa(age)),
			},
			"Gender": {
				N: aws.String(strconv.Itoa(gender)),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("Person"),
	}

	result, err := svc.PutItem(input)
	return result, err
}
