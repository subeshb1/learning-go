package person

import (
	"context"
	"personapi/pkg/model"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetPersonLogic struct {
	Log    logging.Logger
	People PeopleModel
}

type PeopleModel interface {
	Where() ([]model.Person, error)
}

type Person struct {
	Name   string
	Author string
}

func (gl *GetPersonLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {

	people, err := gl.People.Where()
	if err != nil {
		gl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
	}
	res.Body = people
}
