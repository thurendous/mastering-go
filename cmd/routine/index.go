package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("printNumbers %d\n", i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("printLetters %c (ASCII: %d)\n", i, i)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	fmt.Println("This is the main function call")
	// go printLetters()
	time.Sleep(1 * time.Second)
}
