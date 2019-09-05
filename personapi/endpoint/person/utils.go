package person

import (
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Person struct
type Person struct {
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
	UID    string `json:"uid"`
	Name   string `json:"name"`
}

func mapDynoItemToPerson(item map[string]*dynamodb.AttributeValue) Person {
	age, _ := strconv.Atoi(*item["Age"].N)
	gender, _ := strconv.Atoi(*item["Gender"].N)
	person := Person{
		Age:    age,
		Gender: gender,
		UID:    *item["UID"].S,
		Name:   *item["Name"].S,
	}
	return person
}
