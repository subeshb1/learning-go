package main

import "fmt"

func main() {
	i := 1
	b := &i
	*b++
	fmt.Println("initial:", i)
}
