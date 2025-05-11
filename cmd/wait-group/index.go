package main

import (
	"fmt"
	"sync"
)

// add your friends to the wait group before they start their tasks
// 这里的这个wait的地方就是确保所有的goroutine都执行完毕了之后在执行下一步的任务的
var wg sync.WaitGroup

func printMsg(wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	wg.Add(3) // 添加3个goroutine，都是你想要等待的

	go printMsg(&wg, "we have to print this first")
	go printMsg(&wg, "we have to print this second")
	go printMsg(&wg, "we have to print this third")

	wg.Wait()

	fmt.Println("all goroutines finished")
}
