package main

import "fmt"

func main() {
	var m = map[string]int{}

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	fmt.Println(m)

	// delete(m, "Answer")
	// fmt.Println("The value:", m["Answer"])

	// gunakan pengecekan: elem, ok = m[key]
	// v, ok := check(m)
	if v, ok := m["Answer"]; ok {
		fmt.Println(m)
		fmt.Println("The value:", v, "Present?", ok)
	} else {
		fmt.Println(m)
		fmt.Println("Key not found!")
	}
}

func check(m map[string]int) (v int, ok string) {
	if m["Answer"] != 0 {
		return m["Answer"], "yes"
	}
	return -1, "Key not found!"
}
