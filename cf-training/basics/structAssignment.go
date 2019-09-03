// WAP to create a na
package main

import (
	"fmt"
	"strings"
)

type Person struct {
	age                 int
	firstName, lastName string
	gender              string
}

func (p Person) fullName() {
	fmt.Println(strings.ToUpper(p.firstName + " " + p.lastName))
}

func (p Person) birthYear() int {
	return 2019 - p.age
}

func (p *Person) setGender(gender string) {
	switch gender {
	case "M":
		p.gender = "Male"
	case "F":
		p.gender = "Female"
	default:
		p.gender = "Unknown"
	}
}

func main() {
	person1 := Person{firstName: "Subesh", lastName: "Bhandari", age: 22}
	person1.setGender("M")
	fmt.Println(person1.gender)
	fmt.Println(person1.birthYear())
	person1.fullName()
}
