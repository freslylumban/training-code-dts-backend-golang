package main

import "fmt"

func main() {
	var primes []int
	primes = []int{2, 3, 5, 7, 13, 15}

	fmt.Println(primes)
	fmt.Println(primes[2:5])
}
