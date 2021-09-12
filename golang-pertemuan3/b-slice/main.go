package main

import "fmt"

func main() {
	// Array dinamis == SLICE
	var count []int
	count = append(count, 1)
	count = append(count, 2)
	count = append(count, 3)
	fmt.Println(count)

	// Chases
	var buah [4]string = [4]string{"Semangka", "Pisang", "Apel", "Belimbing"}
	var buah2 []string

	for _, value := range buah {
		buah2 = append(buah2, value)
	}

	for _, value := range buah2 {
		fmt.Println(value)
	}

	fmt.Println(len(buah2))
}
