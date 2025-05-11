package main

import "fmt"

func sendData(ch chan<- string) {
	ch <- "1"
	ch <- "2"
	ch <- "3"
	close(ch)
}

func main() {
	// 创建一个整数通道
	dataChannel := make(chan string)

	// 在新的goroutine中运行sendData
	go sendData(dataChannel)

	fmt.Println("channel 接收到的数据:", <-dataChannel)

	// 从通道接收并打印数据
	for value := range dataChannel {
		fmt.Println("接收到的数据:", value)
	}

	fmt.Println("通道已关闭，所有数据接收完毕")

	// 创建一个带缓冲的通道，容量为3
	chn := make(chan int, 3)

	// 发送数据
	chn <- 1
	chn <- 2
	chn <- 3
	// close(chn)

	// 接收数据
	value01 := <-chn
	value02 := <-chn
	value03 := <-chn

	fmt.Println(value01, value02, value03)

	// unbuffered channel
	unbufferedCh := make(chan int)

	// buffered channel
	bufferedCh := make(chan string, 2)

	go func() {
		unbufferedCh <- 1
	}()

	go func() {
		bufferedCh <- "A string"
		bufferedCh <- "B string"
		close(bufferedCh)
	}()

	fmt.Printf("unbufferedCh 接收到的数据: %v\n", <-unbufferedCh)
	fmt.Printf("bufferedCh 接收到的数据1: %v\n", <-bufferedCh)
	fmt.Printf("bufferedCh 接收到的数据2: %v\n", <-bufferedCh)

}
