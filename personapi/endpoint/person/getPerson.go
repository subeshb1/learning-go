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

type DynamoDBConfig interface {
	DynamoDBSession() *dynamodb.DynamoDB
}
type peopleModel interface {
	Where(*dynamodb.DynamoDB) (*dynamodb.QueryOutput, error)
	// Create(name string, age, gender int) (*dynamodb.PutItemOutput, error)
	// Find(name string, age, gender int) (model.Person, error)
}

func (gpl *GetPersonLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	// dynamodb := gpl.Config.DynamoDBSession()
	// people, err := gpl.People.Where(dynamodb)
	// if err != nil {
	// 	gpl.Log.LogErrorf("Could not read data file: %v", err)
	// 	res.HTTPStatus = 400
	// }
	res.Body = "people"
}
