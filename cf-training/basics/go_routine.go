// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

func printUpto(programName string, number int) {
	for i := 1; i <= number; i++ {
		fmt.Println(programName, ":", i)
		time.Sleep(100)
	}
}

func main() {
	go printUpto("First Thread", 5)
	printUpto("Main Thread", 10)
	go printUpto("Second Thread", 10)
	fmt.Scanln()
}
