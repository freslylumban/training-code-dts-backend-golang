package main

import "fmt"

func main() {
	i := 42
	f := 3.142
	g := 0.867 + 0.5i

	fmt.Printf("%T ", i)
	fmt.Printf("%T ", f)
	fmt.Printf("%T ", g)
}
