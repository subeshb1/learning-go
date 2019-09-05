package person

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type AddPersonLogic struct {
	Log    logging.Logger
	People CreatePeople
	Config DynamoDBConfig
}

type CreatePeople interface {
	Create(svc *dynamodb.DynamoDB, name string, age, gender int) (map[string]*dynamodb.AttributeValue, error)
}

func (apl *AddPersonLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, payload *Person) {
	dynamodb := apl.Config.DynamoDBSession()
	person, err := apl.People.Create(dynamodb, payload.Name, payload.Age, payload.Gender)
	if err != nil {
		apl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
		return
	}
	res.HTTPStatus = 201
	res.Body = mapDynoItemToPerson(person)
}
