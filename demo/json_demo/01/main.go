package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := Person{
		Name:    "李四",
		Age:     10,
		address: "四川",
	}
	// 1.将对象转换为json
	jsonArr, _ := json.Marshal(p1)
	fmt.Println(string(jsonArr))

	// 2.反序列化
	jsonStr := `{"Name":"李四","Age":10}`
	var p2 Person
	json.Unmarshal([]byte(jsonStr), &p2) // 必须传指针对象，因为结构体是值类型
	fmt.Printf("%#v \n", p2)

	// 3.结构体标签，序列化/反序列化时起别名
	p3 := Person1{
		Name:    "李四",
		Age:     10,
		Address: "四川",
	}
	jsonArr1, _ := json.Marshal(p3)
	fmt.Println(string(jsonArr1))
	// 反序列化
	jsonStr1 := `{"name":"李四","age":10,"address":"四川"}`
	var p4 Person1
	json.Unmarshal([]byte(jsonStr1), &p4) // 必须传指针对象，因为结构体是值类型
	fmt.Printf("%#v \n", p4)
	// 也可以这样写（区别是上面的是值类型，下面的是指针类型）
	var p5 = &Person1{}
	json.Unmarshal([]byte(jsonStr1), p5)
	fmt.Printf("%#v \n", p5)
}

type Person struct {
	Name    string
	Age     int
	address string // 字段小写开头表示私有，不能被json序列化和反序列化
}

type Person1 struct {
	Name    string `json:"name"` // 也可以定义多个tag,(`json:"name" form:"username"`) key不需要双引号
	Age     int    `json:"age"`
	Address string `json:"address"`
}
