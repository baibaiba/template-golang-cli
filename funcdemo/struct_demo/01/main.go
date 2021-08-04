package main

import "fmt"

// 结构体
func main() {
	// 结构体的实例化方式
	// 1.
	var p1 Person //实例化结构体
	p1.Name = "张三"
	p1.Age = 10
	p1.Sex = "男"
	fmt.Printf("值：%#v,类型：%T \n", p1, p1) // 使用%#v打印结构体更详细
	// 2.new关键字
	p2 := new(Person)
	p2.Name = "张三"
	p2.Age = 10
	p2.Sex = "男"
	fmt.Printf("值：%#v,类型：%T \n", p2, p2) // 类型是结构体指针
	// 在golang中支持结构体指针直接使用，来访问结构体的成员，p2.name = "张三" 在底层是（等效）(*p2).name
	// 3.
	p3 := &Person{}
	p3.Name = "张三"
	p3.Age = 10
	p3.Sex = "男"
	fmt.Printf("值：%#v,类型：%T \n", p3, p3)
	// 4.
	p4 := Person{
		Name: "张三",
		Age:  10,
		Sex:  "男",
	}
	fmt.Printf("值：%#v,类型：%T \n", p4, p4)
	// 5.
	p5 := &Person{
		Name: "张三",
		Age:  10,
		Sex:  "男",
	}
	fmt.Printf("值：%#v,类型：%T \n", p5, p5)
	// 6. 部分属性不赋值，则初始化后属性的值是属性类型的默认值
	p6 := &Person{
		Name: "张三",
	}
	fmt.Printf("值：%#v,类型：%T \n", p6, p6)
	// 7.可以不写key,但是值的类型必须一一对应，并且不能少赋值
	p7 := &Person{
		"张三",
		10,
		"男",
	}
	fmt.Printf("值：%#v,类型：%T \n", p7, p7)
}

// 首字母大写公有，小写私有
type Person struct {
	Name string // 属性后面没有逗号；属性值首字母建议大写，如果属性名小写则在数据解析（如json解析,或将结构体作为请求或访问参数）时无法解析
	Age  int
	Sex  string
}

//------------------------回顾---------------------
// 自定义类型(新类型，打印可看类型)
type myInt int

// 自定义方法类型
type myFunc func(int, int) int

// 类型别名(还是赋值的类型)
type myFloat = float64
