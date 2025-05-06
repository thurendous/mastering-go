package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello, World!")
	}

	f()

	f2 := func(name string) {
		fmt.Println("Hello, " + name + "!")
	}

	f2("Funkei")

	func() {
		fmt.Println("Anonymous function")
	}()
}
