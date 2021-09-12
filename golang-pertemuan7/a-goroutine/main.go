package main

import (
	"fmt"
	"time"
)

func showName(name string) {
	fmt.Println("My name is ", name)
}

func sendEmail(email string) {
	fmt.Println("Ini emailnya: ", email)
}

func main() {
	showName("Fresly")
	go showName("Tren")

	time.Sleep(2 * time.Second)

	sendEmail("freslylu@gmail.com")
	go sendEmail("freslylumban@gmail.com")

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
