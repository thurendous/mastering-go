package common

import "fmt"

// SharedFunction 是一个可以被多个程序共享的函数
func SharedFunction() {
	fmt.Println("This is a shared function")
}
