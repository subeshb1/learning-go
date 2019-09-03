package main

import "fmt"

func add(a, b int, out chan int) {
	out <- a + b
}

func sumAndFormat(a, b int, out chan string) {
	out <- fmt.Sprintf("%d + %d: %d", a, b, a+b)
}

func main() {
	intChannels := make(chan int, 2)
	stringChannels := make(chan string)
	go add(2, 4, intChannels)
	go add(5, 5, intChannels)
	go sumAndFormat(<-intChannels, <-intChannels, stringChannels)
	fmt.Println(<-stringChannels)
}
