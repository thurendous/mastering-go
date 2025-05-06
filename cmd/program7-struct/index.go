package main

import "fmt"

// 定义结构体
type Person struct {
	Name   string
	Age    int
	Job    string
	Salary int
}

type Employee struct {
	Person
	IsProgrammer bool
}

func main() {
	// 创建结构体实例1
	var person1 Person
	person1.Name = "Julia"
	person1.Age = 21
	person1.Job = "Teacher"
	person1.Salary = 4000

	// 创建结构体实例2
	person2 := Person{
		Name:   "John",
		Age:    25,
		Job:    "Engineer",
		Salary: 10000,
	}

	fmt.Println(person1)
	fmt.Println(person2)

	// 访问结构体字段
	fmt.Println(person1.Name)
	fmt.Println(person1.Age)

	// 修改结构体字段
	person1.Age = 21
	fmt.Println(person1)

	// 创建嵌套结构体实例
	employee := Employee{
		Person: Person{
			Name: "Funkei",
			Age:  22,
		},
		IsProgrammer: true,
	}

	fmt.Println(employee)

	// 访问嵌套结构体字段
	fmt.Println(employee.Person.Name, employee.IsProgrammer)
}
