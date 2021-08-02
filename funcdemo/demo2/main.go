package main

import "fmt"

// 匿名函数 返回0
func func1() int {
	var a int
	defer func() {
		a++
	}()
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
	}(x)
	return 5
}
