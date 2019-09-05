package person

import (
	"context"
	"personapi/pkg/model"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type AddPersonLogic struct {
	Log    logging.Logger
	People PeopleModel
	Config DynamoDB
}

type DynamoDB interface {
	DynamoDBSession() *dynamodb.DynamoDB
}
type PeopleModel interface {
	Create(svc *dynamodb.DynamoDB, name string, age, gender int) (map[string]*dynamodb.AttributeValue, error)
}

func (apl *AddPersonLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, payload *model.Person) {
	dynamodb := apl.Config.DynamoDBSession()
	people, err := apl.People.Create(dynamodb, payload.Name, payload.Age, payload.Gender)
	if err != nil {
		apl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
	} else {
		res.HTTPStatus = 201
		res.Body = people
	}
}
