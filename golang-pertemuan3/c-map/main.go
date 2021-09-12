package main

import "fmt"

func main() {
	// MAP
	var ages = map[string]int{} // ages := make(map[string]int)
	ages["alice"] = 31
	ages["beta"] = 37
	ages["charlie"] = 40
	fmt.Println(ages)

	for key, value := range ages {
		fmt.Println("Key: ", key, ", Value: ", value)
	}
}
