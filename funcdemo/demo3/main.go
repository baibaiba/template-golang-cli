package main

import "fmt"

/**
Go语言中目前（1.12）是没有异常机制的，可以使用panic/recover模式处理错误
panic可以在任何地方触发 单recover只有在defer调用的函数中有效
*/
func func1() {
	fmt.Println("func1")
}

func func2() {
	// 这类似Java的try-catch，如果不捕获异常，代码执行终止
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到异常：", err)
		}
	}()
	panic("抛出一个异常")
}

func func3() {
	fmt.Println("func3")
}

func main() {
	func1()
	func2()
	func3()
}
