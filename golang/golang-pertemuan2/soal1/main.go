package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	rand.Seed(10)
// 	fmt.Println("Nomor favorit saya adalah :", rand.Intn(100))
// }

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10
	fmt.Println(rand.Intn(max-min) + min)
}
