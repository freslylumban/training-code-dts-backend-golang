package main

import (
	"fmt"
	"time"
)

func main() {
	// one data
	channels := make(chan int)

	go func() {
		channels <- 72
	}()

	result := <-channels
	fmt.Println(result)

	time.Sleep(1 * time.Second)

	// two data
	channels2 := make(chan int, 2)

	go func() {
		channels2 <- 81
		channels2 <- 36
	}()

	result1 := <-channels2
	result2 := <-channels2
	fmt.Println(result1)
	fmt.Println(result2)

	time.Sleep(1 * time.Second)
}
