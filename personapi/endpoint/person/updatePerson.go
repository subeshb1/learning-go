package person

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type UpdatePersonLogic struct {
	Log    logging.Logger
	People UpdatePeople
	Config DynamoDB
}

type UpdatePeople interface {
	Update(svc *dynamodb.DynamoDB, uid, name string, age, gender int) (map[string]*dynamodb.AttributeValue, error)
}

type UpdateRequest struct {
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
	UID    string `json:"uid"`
	Name   string `json:"name"`
}

func (apl *UpdatePersonLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, payload *UpdateRequest) {
	dynamodb := apl.Config.DynamoDBSession()
	people, err := apl.People.Update(dynamodb, payload.UID, payload.Name, payload.Age, payload.Gender)
	if err != nil {
		apl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
	} else {
		res.HTTPStatus = 201
		res.Body = mapDynoItemToPerson(people)
	}
}
