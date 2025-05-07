package main

import "fmt"

func sendData(ch chan<- int) {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
}

func main() {
	// 创建一个整数通道
	dataChannel := make(chan int)

	// 在新的goroutine中运行sendData
	go sendData(dataChannel)

	fmt.Println("channel 接收到的数据:", <-dataChannel)

	// 从通道接收并打印数据
	for value := range dataChannel {
		fmt.Println("接收到的数据:", value)
	}

	fmt.Println("通道已关闭，所有数据接收完毕")
}
