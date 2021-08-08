package main

import "fmt"

// golang中的类型断言
func main() {
	/**
	类型断言:一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值
	如果我们想要判断空接口中值的类型，那么这个时候就可以使用类型断言，其语法格式：x.(T)
	x：表示类型为interface{}的变量
	T：表示断言x可能是的类型
	*/
	// -----------------------------------------
	// 类型断言的两种方式
	// 1. 变量.（类型）
	var a interface{} = 123
	v, ok := a.(int)
	if ok {
		fmt.Println("a是int类型，断言成功，值：", v)
	} else {
		fmt.Println("a不是int类型，断言失败")
	}

	// 2.（类型.type只能结合switch语句使用）
	MyPrint(1)
	MyPrint("str")
	MyPrint(false)
	MyPrint(1.1)
	// 类型断言也能判断结构体的类型
	structType(Phone{
		Name: "你好",
	})
	structType(123)
}

func MyPrint(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int 类型")
	case string:
		fmt.Println("string 类型")
	case bool:
		fmt.Println("bool 类型")
	case float32:
		fmt.Println("float32 类型")
	case float64:
		fmt.Println("float64 类型")
	default:
		fmt.Println("未知类型")
	}
}

type Phone struct {
	Name string
}

func structType(x interface{}) {
	if _, ok := x.(Phone); ok {
		fmt.Println("Phone类型")
	} else {
		fmt.Println("其它类型")
	}
}
