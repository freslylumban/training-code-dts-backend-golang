package main

import (
	"fmt"
	"time"
)

func main() {
	channels := make(chan string, 2)

	go func() {
		channels <- "Hallo"
	}()

	go func() {
		channels <- "Fresly"
	}()

	result1 := <-channels
	result2 := <-channels
	fmt.Println(result1)
	fmt.Println(result2)

	time.Sleep(1 * time.Second)
}
