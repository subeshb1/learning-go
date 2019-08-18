package main

// CH007 :: LCM of numbers

import "fmt"

func hcf(a, b int) int {
	if a == 0 {
		return b
	}
	return hcf(b%a, a)
}

func lcm(numbers ...int) int {
	result := 1
	for _, number := range numbers {
		result = (number * result) / hcf(result, number)
	}
	return result
}

func main() {
	fmt.Println(lcm(9, 2, 3))       // 	18
	fmt.Println(lcm(9, 2, 3, 1, 4)) // 	36
	fmt.Println(lcm(4, 5, 6, 1))    // 	60
	fmt.Println(lcm(100, 200))      // 	200
	fmt.Println(lcm(50, 21))        // 	1050

}
