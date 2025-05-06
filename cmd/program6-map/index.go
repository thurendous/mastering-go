package main

import "fmt"

func main() {
	ppls := []string{"John", "Jane", "Jim", "Jill"}

	// index, value := - 在每次迭代中，获取当前元素的索引和值
	for index, value := range ppls {
		fmt.Println("index:", index, "value:", value)
	}

	// map
	userInfo := map[string]int{
		"John": 89,
		"Jane": 95,
		"Jim":  78,
		"Jill": 88,
	}
	fmt.Println(userInfo)

	// 遍历map
	for key, value := range userInfo {
		fmt.Println("key:", key, "value:", value)
	}

	// 添加元素
	userInfo["Alex"] = 60
	fmt.Println(userInfo)

	// 修改元素
	userInfo["John"] = 90
	fmt.Println(userInfo)

	// 删除元素
	delete(userInfo, "Alex")
	fmt.Println(userInfo)

	// 检查元素是否存在
	value, ok := userInfo["Alex"]
	if ok {
		fmt.Println("Alex's score:", value)
	} else {
		fmt.Println("Alex's score is not in the map")
	}

}
