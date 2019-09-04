package person

import (
	"context"
	"personapi/pkg/model"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type AddPersonLogic struct {
	Log    logging.Logger
	People PeopleModel
}

type PeopleModel interface {
	Create(name string, age, gender int) (model.Person, error)
}

type Person struct {
	Name   string
	Author string
}

func (apl *AddPersonLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, payload *model.Person) {

	people, err := apl.People.Create(payload.Name, payload.Age, payload.Gender)
	if err != nil {
		apl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
	} else {
		res.HTTPStatus = 201
		res.Body = people
	}
}
