package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

type Employee struct {
	name     string
	address  string
	position string
}

func main() {
	// ARRAY
	// var score [3]int = [3]int{40, 30, 60}
	// fmt.Println(score)

	// for i := 0; i < len(score); i++ {
	// 	fmt.Println(score[i])
	// }

	// for _, value := range score {
	// 	fmt.Println(value)
	// }

	// Array dinamis == SLICE
	// var count []int
	// count = append(count, 1)
	// count = append(count, 2)
	// count = append(count, 3)
	// fmt.Println(count)

	// Chases
	// var buah [4]string = [4]string{"Semangka", "Pisang", "Apel", "Belimbing"}
	// var buah2 []string

	// for _, value := range buah {
	// 	buah2 = append(buah2, value)
	// }

	// for _, value := range buah2 {
	// 	fmt.Println(value)
	// }

	// fmt.Println(len(buah2))

	// MAP
	// var ages = map[string]int{} // ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["beta"] = 37
	// ages["charlie"] = 40
	// fmt.Println(ages)

	// for key, value := range ages {
	// 	fmt.Println("Key: ", key, ", Value: ", value)
	// }

	// STRUCT
	var arrPerson []Person
	arrPerson = append(arrPerson, Person{
		name:    "Alpha",
		age:     23,
		address: "Bekasi",
	})

	var person Person
	person.name = "Beta"
	person.age = 23
	person.address = "Jakarta"
	fmt.Println(arrPerson)
	fmt.Println(person)

	// Chases
	var employee Employee
	employee.name = "Julio"
	employee.address = "Jakarta"
	employee.position = "Staff"

	fmt.Println(employee)

	var arrEmployee []Employee
	arrEmployee = append(arrEmployee, Employee{
		name:     "Tina",
		address:  "Bekasi",
		position: "HRD",
	}, Employee{
		name:     "Chris",
		address:  "Bogor",
		position: "Staff",
	})
	fmt.Println(arrEmployee)
}
