package main

import "fmt"

// 接口体值接收者和指针接收者实现接口的区别
func main() {
	var d = Dog{}
	var d1 Animaler = d
	d1.SetName("大佬狗")
	fmt.Println(d1.GetName())

	var d2 = &Dog{}
	var d3 Animaler = d2
	d3.SetName("大佬狗")
	fmt.Println(d3.GetName())

	// 如果方法接收者是指针类型，那么让值类型Dog实现Animaler会报错，所以只能使用指针类型实现接口（Dog does not implement Animaler (GetName method has pointer receiver)）
	// 如果方法接收者是值类型，那么让值类型或者指针类型实现接口都是可以的
}

type Animaler interface {
	SetName(string)
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(x string) { // 如果方法接收者是值类型是不能修改结构体内属性的
	d.Name = x
}

func (d *Dog) GetName() string {
	return d.Name
}
