package main

import "fmt"

func doSomething(s string, callback func()) {
	fmt.Println(s)
	callback()
}

func printALine() {
	fmt.Println("new line")
}

func main() {
	doSomething("Hello, World!", printALine)
}
