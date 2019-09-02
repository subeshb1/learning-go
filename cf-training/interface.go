package main

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	numSeats() int
	brand() string
}

type Car struct {
	seats     int
	brandName string
}

func (c Car) brand() string {
	return c.brandName
}
func (c Car) numSeats() int {
	return c.seats
}

type Bike struct {
	seats     int
	brandName string
}

func (b Bike) brand() string {
	return b.brandName
}
func (b Bike) numSeats() int {
	return b.seats
}

func getType(v interface {
	numSeats() int
	brand() string
}) string {
	return reflect.TypeOf(v).String()
}

func main() {
	var v Vehicle = Car{brandName: "Bugati", seats: 4}
	fmt.Println(v.brand())
	fmt.Println(v.numSeats())
	fmt.Println(getType(v))
	v = Bike{brandName: "Ducati", seats: 2}
	fmt.Println(v.brand())
	fmt.Println(v.numSeats())
	fmt.Println(getType(v))
	fmt.Println(getType(1))
}
