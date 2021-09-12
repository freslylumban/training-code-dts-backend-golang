package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)

	go func() {
		ch <- "Hello"
	}()

	go func() {
		ch <- "Goroutine"
	}()

	messageFirst := <-ch
	messageLast := <-ch
	fmt.Println(messageFirst)
	fmt.Println(messageLast)
}
