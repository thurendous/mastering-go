package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) GetInfo() string {
	return "Name: " + u.Name + ", Age: " + strconv.Itoa(u.Age)
}

func main() {
	user := User{
		Name: "Funkei",
		Age:  20,
	}

	fmt.Println(user.GetName())
	fmt.Println(user.GetAge())
	fmt.Println(user.GetInfo())
}
