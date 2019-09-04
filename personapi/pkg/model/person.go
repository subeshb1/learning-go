package model

// Person Model
type Person struct {
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
	Name   string `json:"name"`
}

func (p Person) Update() (Person, error) {

	return Person{Age: 10, Gender: 1, Name: "Subesh"}, nil
}

// People Model
type People struct {
}

func (p People) Where() ([]Person, error) {

	return []Person{Person{Age: 10, Gender: 1, Name: "Subesh"}}, nil
}

func (p People) Find() (Person, error) {

	return Person{Age: 10, Gender: 1, Name: "Subesh"}, nil
}

func (p People) Create(name string, age, gender int) (Person, error) {

	return Person{Age: 10, Gender: 1, Name: "Subesh"}, nil
}
