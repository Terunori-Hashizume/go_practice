package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := struct {
		firstName string
		lastName  string
		id        int
	}{
		firstName: "Jane",
		lastName:  "Smith",
		id:        1,
	}
	emp2 := Employee{
		firstName: "John",
		lastName:  "Doe",
		id:        2,
	}
	var emp3 struct {
		firstName string
		lastName  string
		id        int
	}
	emp3.firstName = "Jim"
	emp3.lastName = "Beam"
	emp3.id = 3

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
