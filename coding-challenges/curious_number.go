package main

// CH005 :: Curious Number

import (
	"fmt"
	"math"
)

func isCuriousNumber(number int) bool {
	numberLength := int(math.Log10(float64(number)) + 1)
	return (number*number)%int(math.Pow10(numberLength)) == number
}

func main() {
	fmt.Println(isCuriousNumber(25))        // true
	fmt.Println(isCuriousNumber(76))        // true
	fmt.Println(isCuriousNumber(212890625)) // true
	fmt.Println(isCuriousNumber(1))         // true
	fmt.Println(isCuriousNumber(30))        // false
	fmt.Println(isCuriousNumber(10))        // false
	fmt.Println(isCuriousNumber(229348))    // false

}
