package main

import "fmt"

func main() {
	// ARRAY
	var score [3]int = [3]int{40, 30, 60}
	fmt.Println(score)

	for i := 0; i < len(score); i++ {
		fmt.Println(score[i])
	}

	for _, value := range score {
		fmt.Println(value)
	}
}
