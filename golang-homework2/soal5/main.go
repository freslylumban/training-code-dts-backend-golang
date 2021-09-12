package main

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	// a := add(1, 2)
	// a := add(1, 3, 7)
	// a := add([]int{1, 2})
	// a := add([]int{1, 3, 7}...)
	// fmt.Println(a)

	// SOAL6
	// var x = nil
	// var y interface{} = nil
	// var z string = nil
	// var a error = nil

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(a)

	//SOAL7
	// s := make([]int)
	// t := make([]int, 0)
	// u := make([]int, 5, 10)
	// v := []int{1, 2, 3, 4, 5}

	// fmt.Println(s)
	// fmt.Println(t)
	// fmt.Println(u)
	// fmt.Println(v)
}
