package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("end") // this execution will be delayed to the end of the function
	fmt.Println("middle")
	defer fmt.Println("middle2")
	defer fmt.Println("middle3")
}
