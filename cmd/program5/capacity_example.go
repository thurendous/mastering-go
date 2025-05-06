package main

import (
	"fmt"
)

func ShowCapacityExample() {
	// 创建一个类似于"10座位的公交车，当前坐了3人"的切片
	// 长度=3（当前元素），容量=10（底层数组大小）
	bus := make([]string, 3, 10)
	bus[0] = "张三" // 第1个乘客
	bus[1] = "李四" // 第2个乘客
	bus[2] = "王五" // 第3个乘客

	fmt.Println("初始状态:")
	fmt.Printf("bus = %v\n", bus)
	fmt.Printf("乘客数量(长度) = %d\n", len(bus))
	fmt.Printf("座位数量(容量) = %d\n", cap(bus))

	// 添加乘客，但未超过容量
	fmt.Println("\n添加4个乘客:")
	bus = append(bus, "赵六", "钱七", "孙八", "周九")
	fmt.Printf("bus = %v\n", bus)
	fmt.Printf("乘客数量(长度) = %d\n", len(bus))
	fmt.Printf("座位数量(容量) = %d\n", cap(bus))
	fmt.Println("底层数组地址:", &bus[0]) // 打印底层数组的地址

	// 继续添加乘客，超过原容量
	fmt.Println("\n继续添加4个乘客(超过10个座位):")
	bus = append(bus, "吴十", "郑十一", "王十二", "刘十三")
	fmt.Printf("bus = %v\n", bus)
	fmt.Printf("乘客数量(长度) = %d\n", len(bus))
	fmt.Printf("座位数量(容量) = %d\n", cap(bus)) // 容量会翻倍
	fmt.Println("底层数组地址:", &bus[0])         // 地址变了，说明创建了新数组
}
