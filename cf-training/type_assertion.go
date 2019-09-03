package main

import (
	"fmt"
	"strings"
	"time"
)

func returnMessageToChannel(program string, channel chan string) {
	channel <- strings.ToUpper(program)
}

func main() {
	messages := make(chan string, 2)

	go returnMessageToChannel("First Program", messages)
	time.Sleep(10000)
	go returnMessageToChannel("Second Program", messages)

	first := <-messages
	second := <-messages
	fmt.Println(first, " ran First")
	fmt.Println(second, " ran Second")
}
