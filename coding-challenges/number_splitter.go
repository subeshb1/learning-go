// CH003 :: Number Splitter
package main

import "fmt"

func numberSplitter(number, partition int) []int {
	var partitionArray []int
	if partition <= number && partition > 0 {
		partitionArray = make([]int, partition)
		quotient := number / partition
		if number%partition == 0 {
			for index := range partitionArray {
				partitionArray[index] = quotient
			}
		} else {
			remainder := number - (partition-2)*quotient
			second := remainder / 2
			first := second + remainder%2
			partitionArray[0] = first
			partitionArray[1] = second
			for index := 2; index < partition; index++ {
				partitionArray[index] = quotient
			}
		}
	}
	return partitionArray
}

func main() {
	fmt.Println(numberSplitter(10, 5))   // [2 2 2 2 2]
	fmt.Println(numberSplitter(10, 5))   // [2 2 2 2 2]
	fmt.Println(numberSplitter(10, 4))   // [3 3 2 2]
	fmt.Println(numberSplitter(11, 20))  // []
	fmt.Println(numberSplitter(100, 15)) // [11 11 6 6 6 6 6 6 6 6 6 6 6 6 6]
	fmt.Println(numberSplitter(100, 13)) // [12 11 7 7 7 7 7 7 7 7 7 7 7]
	fmt.Println(numberSplitter(9, 5))    // [3 3 1 1 1]
	fmt.Println(numberSplitter(9, 1))    // [9]
	fmt.Println(numberSplitter(9, 9))    // [1 1 1 1 1 1 1 1 1]
	fmt.Println(numberSplitter(9, -1))   // []

}
