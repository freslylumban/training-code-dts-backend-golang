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
