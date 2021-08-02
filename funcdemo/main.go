package main

import (
	"fmt"
)

/**
1.全局变量：
	常驻内存、污染全局
2.局部变量：
	不常驻内存、不污染全局
3.闭包
	常驻内存、不污染全局
*/

// 闭包：函数里面嵌套一个函数，最后返回里面的函数
func adder1() func() int {
	i := 10 // 常驻内存 ，不污染全局
	return func() int {
		return i + 10
	}
}

// 全局变量修改值改变
func adder2() func(y int) int {
	i := 10 // 常驻内存 ，不污染全局
	return func(y int) int {
		i += y
		return i
	}
}

func main() {
	fnc := adder1()
	fmt.Println(fnc())
	fmt.Println(fnc())
	fmt.Println(fnc())
	fmt.Println("------------------------------")
	fnc1 := adder2()
	fmt.Println(fnc1(10))
	fmt.Println(fnc1(10))
	fmt.Println(fnc1(10))
}
