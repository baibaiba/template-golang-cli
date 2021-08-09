package main

import "fmt"

/**
defer先注册后执行
defer注册要延迟执行的函数时该函数所有的参数需要确定其值
执行顺序如下：
fmt.Println("开始")
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
fmt.Println("结束")
结果打印：
开始
结束
3
2
1
*/

// 匿名函数 返回0
func func1() int {
	var a int
	defer func() {
		a++
	}() // defer后面的函数必须是匿名自执行函数
	return a
}

// 命名函数返回1
func func2() (a int) {
	defer func() {
		a++
	}()
	return a // 返回操作以后的值
}

func main() {
	fmt.Println(func1())
	fmt.Println(func2())
	fmt.Println("--------------------")
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

/**
测试方法
*/
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	// 先将5赋值给返回值x,然后执行x++
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		// 此时x = 0
		x++
	}(x) // defer注册要延迟执行的函数时该函数所有的参数需要确定其值
	// 这里定义了一个函数，并且传参x,所以刚开始参数值就需要确定 x=0
	return 5
}
