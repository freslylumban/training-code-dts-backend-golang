package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	// fmt.Println("Halo nama saya Fresly")

	// var name string
	// name = "Fresly"
	// fmt.Println(name)

	// var age int
	// age = 27
	// fmt.Println("Umur saya", age)

	// var isOpen bool
	// isOpen = false
	// fmt.Println(isOpen)

	// // Shorthand var
	// count := 20
	// artist := "Blackpink"

	// if artist == "Blackpink" && count == 20 {
	// 	fmt.Println("Blackpink and count 20")
	// } else {
	// 	fmt.Println("This is not Blackpink")
	// }

	number := 300
	pointerNumb := &number
	fmt.Println(pointerNumb)
	fmt.Println(*pointerNumb)
}
