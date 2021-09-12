package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i

	fmt.Println(i, j, *p)

	*p = 21

	fmt.Println(i)

	p = &j

	j = (*p / 37)

	fmt.Println(j)
}
