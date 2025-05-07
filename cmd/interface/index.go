package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func announce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	announce(dog)
	announce(cat)
}
