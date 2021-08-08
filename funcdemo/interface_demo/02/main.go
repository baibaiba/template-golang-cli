package main

import "fmt"

// golang中的interface
func main() {
	// interface{}表示一个空接口，可以表示任意类型
	// 举例：
	arr := [](interface{}){123, "str", false}
	fmt.Println(arr)

	map1 := make(map[string]interface{})
	map1["a"] = 1
	map1["b"] = "str"
	map1["c"] = false
	fmt.Println(map1)

	func1(1)
	func1("str")
	func1(1.23)
	func1(true)
}

// 空接口为函数参数
func func1(x interface{}) {
	fmt.Println(x)
}
