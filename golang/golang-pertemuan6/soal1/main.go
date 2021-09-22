package main

import "fmt"

func main() {
	var a [2]string = [2]string{"Go", "C"}
	fmt.Println(a[0])
	fmt.Println(a[1])

	var primes [6]int
	primes = [6]int{2, 3, 5, 7, 13, 15}
	fmt.Println(primes)
}
