package main

import (
	"fmt"
	"mastering-go/pkg/common"
)

func main() {
	fmt.Println("This is program 2")
	common.SharedFunction()

	// if, else if, else
	// if condition { <code> } else if condition { <code> } else { <code> }

	num := 5
	if num > 5 {
		fmt.Println("num is greater than 5")
	} else if num < 5 {
		fmt.Println("num is less than 5")
	} else {
		fmt.Println("num is equal to 5")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
