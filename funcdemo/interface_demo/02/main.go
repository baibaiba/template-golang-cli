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

	// interface{}可能出现的问题
	map2 := make(map[string]interface{})
	map2["a"] = 1
	map2["hobby"] = []string{"吃饭", "睡觉", "打豆豆"}
	map2["address"] = Address{
		Name:  "李四",
		Phone: 15299999999,
	}
	fmt.Println(map2)
	// 这个时候我们获取其中的一个爱好和收货人名称?
	// 不经思索的获取方式
	// hobby := map2["hobby"][0] // 错误 (type interface {} does not support indexing)
	// fmt.Println(hobby)
	// name := map2["address"].Name // 错误 (type interface {} is interface with no methods)
	// fmt.Println(name)
	// 正确方式如下：结合类型断言获取值
	hobby, _ := map2["hobby"].([]string) // 如果断言成功，会将值赋值给第一个返回结果
	fmt.Println(hobby)
	fmt.Println(hobby[0])

	address, _ := map2["address"].(Address)
	fmt.Println(address)
	fmt.Println(address.Name, address.Phone)
}

type Address struct {
	Name  string
	Phone int
}

// 空接口为函数参数
func func1(x interface{}) {
	fmt.Println(x)
}
