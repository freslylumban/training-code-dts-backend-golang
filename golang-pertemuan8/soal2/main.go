package main

import "fmt"

func myfunc(ch chan int) {
	fmt.Println(234 + <-ch)
}

func main() {
	fmt.Println("Start Main method")

	ch := make(chan int)
	go myfunc(ch)
	ch <- 73
	fmt.Println("End main method")
}
