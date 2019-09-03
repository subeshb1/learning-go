// WAP to create a funtion, personBuilder(name, age), that builds and returns a Person struct. The function takes in Name and Age,
// and builds the Person struct, but it has birthYear attribute.
// Also, add error handling in the function, that checks if the age is negative.
// Person {name: “barun”, birthYear: 1989 }

package main

import (
	"errors"
	"fmt"
)

type Person struct {
	birthYear int
	name      string
}

func personBuilder(name string, age int) (Person, error) {
	if age < 0 {
		return Person{}, errors.New("Age can't be negative")
	}
	return Person{name: name, birthYear: 2019 - age}, nil

}

func main() {
	person1, error1 := personBuilder("Subesh", 21)
	if error1 != nil {
		fmt.Println(error1)
	} else {
		fmt.Println(person1)
	}
	person2, error2 := personBuilder("Subesh", -1)
	if error2 != nil {
		fmt.Println(error2)
	} else {
		fmt.Println(person2)
	}
}
