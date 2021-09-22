package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("NAMA")
	fmt.Println("Hello World,", name)
}
