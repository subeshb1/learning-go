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
}

func (p Person) All(svc *dynamodb.DynamoDB) ([]map[string]*dynamodb.AttributeValue, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Person"),
	}

	result, err := svc.Scan(input)
	return result.Items, err
}

func (p Person) Find(svc *dynamodb.DynamoDB, uid string) (map[string]*dynamodb.AttributeValue, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UID": {
				S: aws.String(uid),
			},
		},
		TableName: aws.String("Person"),
	}

	result, err := svc.GetItem(input)
	return result.Item, err
}

func (p Person) Create(svc *dynamodb.DynamoDB, name string, age, gender int) (map[string]*dynamodb.AttributeValue, error) {
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

	_, err := svc.PutItem(input)
	if err != nil {
		return nil, err
	}
	person, err := p.Find(svc, uid)
	return person, err
}

func (p Person) Update(svc *dynamodb.DynamoDB, uid, name string, age, gender int) (map[string]*dynamodb.AttributeValue, error) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("Name"),
			"#A": aws.String("Age"),
			"#G": aws.String("Gender"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(name),
			},
			":a": {
				N: aws.String(strconv.Itoa(age)),
			},
			":g": {
				N: aws.String(strconv.Itoa(gender)),
			},
			":u": {
				S: aws.String(uid),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"UID": {
				S: aws.String(uid),
			},
		},
		ConditionExpression: aws.String("UID = :u"),
		ReturnValues:        aws.String("ALL_NEW"),
		TableName:           aws.String("Person"),
		UpdateExpression:    aws.String("SET #N = :n, #A = :a, #G = :g"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		return nil, err
	}
	person, err := p.Find(svc, uid)
	return person, err
}
