package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, start!")
	var numbers [5]int

	var peoples [3]string

	peoples[0] = "张三"
	peoples[1] = "李四"
	peoples[2] = "王五"

	// first time print
	fmt.Println(numbers)
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	// second time print
	fmt.Println(numbers)

	// 调用printNumbers函数
	printNumbers(numbers)

	// 调用printPeoples函数
	printPeoples(peoples)
	fmt.Println(len(peoples))

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	for i := 0; i < len(arr1); i++ {
		fmt.Println("arr1's element:", arr1[i])
	}

	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

}

func printNumbers(numbers [5]int) {
	fmt.Println(numbers)
}

func printPeoples(peoples [3]string) {
	fmt.Println(peoples)
}
