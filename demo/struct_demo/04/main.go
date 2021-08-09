package main

import "fmt"

// 结构体的嵌套
// 结构体的字段类型可以是: 切片、基本数据类型、map以及结构体
// 如果结构体的字段类型是：指针、切片、map的默认值是nil，即还没有分配内存空间，如果需要使用这些类型的字段需要先分配内存空间
// 如果结构体中需要组合其他结构体，那么建议采用指针的方式去声明，否则会出现更新丢失问题。
func main() {
	var p1 Person
	p1.Name = "张三"
	p1.Age = 10
	p1.Sex = "男"
	p1.Address.Address = "四川省成都市高新区益州大道"
	p1.Address.City = "成都"
	p1.Address.Name = "dove"
	p1.Address.Phone = "181xxxxxxxx"
	fmt.Printf("%#v \n", p1)
	// 设置值时如果结构体与嵌套类型中字段名一样，会先设置结构体中的字段
	var p2 Person
	p2.Name = "李四"
	fmt.Printf("%#v \n", p2)
	// 如果嵌套类型中字段相同，设置值时需要指定类型，不然会报错
	// 同理，获取值时先获取结构体的字段值，如果结构体中没有这个字段，会在嵌套类型中查找
	fmt.Printf("%v \n", p1.Name)
}

type Person struct {
	Name    string
	Age     int
	Sex     string
	Address Address // 类型可省略，
}

type Address struct {
	Address string
	City    string
	Name    string
	Phone   string
}
