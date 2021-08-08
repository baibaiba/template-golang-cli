package main

import (
	"fmt"

	"github.com/gitpod/mycli/funcdemo/go_mod_demo/tools"
)

// golang中init方法的执行顺序
func main() {
	// 定义：在go语言程序执行时导入包语句会自动触发包内部init函数的调用。
	// 注意：init函数没有参数也没有返回值。init函数在程序运行时自动被调用，不能在代码中主动调用
	add := tools.Add(1, 2)
	fmt.Println(add)
	fmt.Println("main方法执行")
}

func init() {
	fmt.Println("main init")
}
