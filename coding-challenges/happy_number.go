package main

// CH004 :: Happy Number

import (
	"fmt"
	"math"
)

func squaredDigitSum(number int) int {
	sum := 0
	for number != 0 {
		sum += int(math.Pow(float64(number%10), 2))
		number = number / 10
	}
	return sum
}

func isHappyNumber(number int) bool {
	memo := make(map[int]int)
	memo[number] = 1
	for nextNumber := squaredDigitSum(number); nextNumber != 1; nextNumber = squaredDigitSum(nextNumber) {
		if _, present := memo[nextNumber]; present {
			return false
		}
		memo[nextNumber] = 1
	}
	return true
}

func main() {
	fmt.Println(isHappyNumber(7))  // true
	fmt.Println(isHappyNumber(10)) // true
	fmt.Println(isHappyNumber(2))  // false
	fmt.Println(isHappyNumber(11)) // false
	fmt.Println(isHappyNumber(13)) // true
	fmt.Println(isHappyNumber(19)) // true
	fmt.Println(isHappyNumber(9))  // false

}
