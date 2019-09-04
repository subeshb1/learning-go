package person

import (
	"context"
	"personapi/pkg/model"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetPersonLogic struct {
	Log    logging.Logger
	People peopleModel
}

type peopleModel interface {
	Where() ([]model.Person, error)
}

func (gpl *GetPersonLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {

	people, err := gpl.People.Where()
	if err != nil {
		gpl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
	}
	res.Body = people
}
