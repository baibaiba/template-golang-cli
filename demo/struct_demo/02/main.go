package main

import "fmt"

/**
值类型:数组、基本数据类型、结构体
引用类型：map、slice
*/
// 结构体是值类型
func main() {
	p1 := Person{
		"张三",
		10,
		"男",
	}

	p2 := p1
	p2.Name = "李四"
	fmt.Printf("p1值：%#v,类型：%T \n", p1, p1)
	fmt.Printf("p2值：%#v,类型：%T \n", p2, p2)

	// 这是引用类型
	p3 := &Person{
		"张三",
		10,
		"男",
	}
	p4 := p3
	p4.Name = "李四"
	fmt.Printf("p3值：%#v,类型：%T \n", p3, p3)
	fmt.Printf("p4值：%#v,类型：%T \n", p4, p4)
}

type Person struct {
	Name string
	Age  int
	Sex  string
}
