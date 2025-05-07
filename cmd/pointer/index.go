package main

import "fmt"

func main() {
	address := "123 Main St"

	fmt.Println(address)
	fmt.Println(&address)

	var pointer *string = &address
	fmt.Println("memory address: ", pointer)
	fmt.Println("value: ", *pointer)
}
