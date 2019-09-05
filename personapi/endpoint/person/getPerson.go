package person

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetPersonLogic struct {
	Log    logging.Logger
	People AllPeople
	Config DynamoDBConfig
}

type AllPeople interface {
	All(*dynamodb.DynamoDB) ([]map[string]*dynamodb.AttributeValue, error)
}

func (gpl *GetPersonLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	dynamodb := gpl.Config.DynamoDBSession()
	people, err := gpl.People.All(dynamodb)
	if err != nil {
		gpl.Log.LogErrorf("Failed to fetch data: %v", err)
		res.HTTPStatus = 500
		return
	}
	res.HTTPStatus = 200
	res.Body = mapDynamoDBItemsToPeople(people)
}
