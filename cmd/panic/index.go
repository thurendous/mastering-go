package main

import "strconv"

func employee(name *string, age int) {
	if age > 65 {
		panic("Employee" + *name + " is too old: " + strconv.Itoa(age))
	}
}

func main() {
	empName := "John"
	age := 70
	employee(&empName, age)
}
