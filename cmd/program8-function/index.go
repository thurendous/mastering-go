package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	// call function
	sayHello()
	greet("Funkei")
	greet2(100)
	showNumbers(1, 3, 5, 7, 9, 11)

	// call function, has return value
	result := add(1, 2)
	fmt.Println(result)

	// call function, has multiple return value
	divide, remainder := divide(10, 3)
	fmt.Println("divided result: ", divide, "remainder: ", remainder)

	// call function, has variable parameter
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)

	// call function, has anonymous function
	addition, subtraction := calculate(10, 5)
	fmt.Println("addition: ", addition, "subtraction: ", subtraction)

}

// define a function
func sayHello() {
	fmt.Println("Hello, World!")
}

// define a function, has parameter
func greet(name string) {
	fmt.Println("Hello, " + name + "!")
}

func greet2(num int) {
	fmt.Println("Hello, " + strconv.Itoa(num) + "!")
}

// define a function, has return value
func add(a int, b int) int {
	return a + b
}

// define a function, has multiple return value
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// define a function, has variable parameter
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func showNumbers(numbers ...int) {
	fmt.Println(numbers)
}

// define a function, has anonymous function
func calculate(a int, b int) (int, int) {
	return a + b, a - b
}
