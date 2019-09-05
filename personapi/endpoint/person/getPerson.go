package person

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetPersonLogic struct {
	Log    logging.Logger
	People peopleModel
	Config DynamoDBConfig
}

type peopleModel interface {
	All(*dynamodb.DynamoDB) ([]map[string]*dynamodb.AttributeValue, error)
}

func (gpl *GetPersonLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	dynamodb := gpl.Config.DynamoDBSession()
	people, err := gpl.People.All(dynamodb)
	if err != nil {
		gpl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
		return
	}
	res.HTTPStatus = 200
	res.Body = mapDynoItemsToPeople(people)
}
