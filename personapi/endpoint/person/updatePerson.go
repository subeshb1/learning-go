package person

// import (
// 	"context"
// 	"personapi/pkg/model"

// 	"github.com/graniticio/granitic/v2/logging"
// 	"github.com/graniticio/granitic/v2/ws"
// )

// type UpdatePersonLogic struct {
// 	Log    logging.Logger
// 	People peopleModel
// }

// type peopleModel interface {
// 	Create(name string, age, gender int) (model.Person, error)
// }

// type UpdateRequest struct {
// 	id     int    `json:"age"`
// 	Age    int    `json:"age"`
// 	Gender int    `json:"gender"`
// 	Name   string `json:"name"`
// }

// func (apl *UpdatePersonLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, payload *model.Person) {

// 	people, err := apl.People.Create(payload.Name, payload.Age, payload.Gender)
// 	if err != nil {
// 		apl.Log.LogErrorf("Could not read data file: %v", err)
// 		res.HTTPStatus = 400
// 	} else {
// 		res.HTTPStatus = 201
// 		res.Body = people
// 	}
// }
