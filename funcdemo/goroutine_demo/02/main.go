package main

import (
	"fmt"
	"runtime"
)

// 设置golang并行运行时占用的cpu数量（了解）
func main() {

	num := runtime.NumCPU()
	fmt.Println("运行时占用cpu数量：", num)

	// 设置运行时使用的cpu核数
	runtime.GOMAXPROCS(num)

	// golang 1.5之前，默认使用的是单核心运行。1.5版本之后，默认使用全部的cpu逻辑核心数
}
