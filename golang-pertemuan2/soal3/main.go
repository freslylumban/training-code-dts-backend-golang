package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	x, y = 3, 4

	xFl := float64(x)
	yFl := float64(y)

	total := (xFl*xFl + yFl*yFl)

	sqrtTotal := math.Sqrt(total)

	fmt.Println("test1", x, y, total, sqrtTotal)

	// var z uint
	z := uint(sqrtTotal)

	fmt.Println("test2", x, y, total, sqrtTotal, z)
	fmt.Println("test3", x, y, total, sqrtTotal, z)
}
