// package declaration, indicating the package to which the file belongs
// it plays an important role in managing dependencies and code organization
// The main package is special, it is the entry point of the program
package main

// import the packages the program rely on
// fmt is for formatiting and printing
import (
	"fmt"
	"strings"
)

// main function, the entry point of the program
// function is the entry point for executable program in Go.
// it is where the program begins

// in Golang, string should be enclosed in double quotes
var name = "John"          // variable declaration -> inferred type is string. This let the golang compiler infer the type of the variable
var name3 string = "John3" // variable declaration -> explicit type is string
var num = 1002

const constantValue = 123 // this vlaue is a constant one and not changeable
const (
	a = 1
	b = 2
	c = 3
)

func main() {
	fmt.Println("Hello, World!")

	name2 := "John2" // variable declaration
	number1 := 1001

	// boolean
	isGolangPL := true
	isHtmlPL := false

	// float number
	pi := 3.14
	// integer number
	multiple := 3

	// string
	// this can make the string multi-line
	message := `Hello, 
	World! 
	My name 
	is HuXin`
	message2 := "HeLlo"
	msg := "one"
	msg2 := "One"

	// print the var value and the type of the var
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(number1)
	fmt.Println(name3)
	fmt.Println(num)
	fmt.Println(constantValue)
	fmt.Println(a, b, c)
	fmt.Println(isGolangPL)
	fmt.Println(isHtmlPL)

	fmt.Println(int(pi) * multiple)
	fmt.Println(pi * float64(multiple))
	fmt.Println(message)

	fmt.Printf("%c\n", message2[0])
	fmt.Println(len(message2))
	fmt.Printf("%v\n", message2[2])

	fmt.Println(strings.Compare(msg, msg2))

	fmt.Println(strings.Contains(msg2, "O"))
	fmt.Println(strings.Contains(msg2, "o"))

	fmt.Println(strings.Count(msg2, "o"))
	fmt.Println(strings.Count(msg2, "O"))
	fmt.Println(strings.ToLower(msg2))
	fmt.Println(strings.ToUpper(msg2))

	fmt.Println("================================================")

	// print the var value and the type of the var
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of name2: %T\n", name2)
	fmt.Printf("Type of name3: %T\n", name3)
	fmt.Printf("Type of number1: %T\n", number1)
	fmt.Printf("Type of constantValue: %T\n", constantValue)

	fmt.Printf("Type of isGolangPL: %T\n", isGolangPL)
	fmt.Printf("Type of isHtmlPL: %T\n", isHtmlPL)
	fmt.Printf("Type of pi: %T\n", pi)
	fmt.Printf("Type of mutiple: %T\n", multiple)
	fmt.Printf("Type of message: %T\n", message2[2])
}
