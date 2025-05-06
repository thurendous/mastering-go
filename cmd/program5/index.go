package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Println(myslice1)
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1)) // capacity

	myslice1 = append(myslice1, 1, 2, 3)
	fmt.Println(myslice1)

	// make()
	// make func can create a zeroed array and a lice referencing an array.
	// This is a great way to create a dynamically sized array.
	// parameters: type, length, capacity
	slice := make([]string, 3, 10) // 创建一个长度为3、容量为10的切片
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 容量示例
	fmt.Println("\n容量示例:")
	// 创建一个长度为3、容量为5的切片
	s1 := make([]int, 3, 5)
	fmt.Println("初始状态:")
	fmt.Printf("s1=%v, 长度=%d, 容量=%d\n", s1, len(s1), cap(s1))

	// 添加元素不超过容量
	s1 = append(s1, 10, 20)
	fmt.Println("\n添加2个元素后(未超过容量):")
	fmt.Printf("s1=%v, 长度=%d, 容量=%d\n", s1, len(s1), cap(s1))

	// 继续添加元素,超过容量
	s1 = append(s1, 30)
	fmt.Println("\n再添加1个元素后(超过原容量):")
	fmt.Printf("s1=%v, 长度=%d, 容量=%d\n", s1, len(s1), cap(s1))

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2[:])
	fmt.Println(arr2[0:2])
	fmt.Println(arr2[2:4])
	fmt.Println(arr2[1:5])

	// 数组遍历
	for i := 0; i < len(arr2); i++ {
		fmt.Println("arr2's element:", arr2[i])
	}

	// 运行公交车容量示例
	fmt.Println("\n运行公交车容量示例:")
	ShowCapacityExample()
}
