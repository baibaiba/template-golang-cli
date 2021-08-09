package main

import (
	"fmt"
	"time"
)

// goroutine 中出现panic解决
func main() {
	go func1()
	go func2() // 如果在主线程中其中一个goroutine方法报错，其它协程都不会执行
	// 解决方法：defer + recover捕获异常
	time.Sleep(time.Second)
}

func func1() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello word")
	}
}

func func2() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("func2 发生错误", error)
		}
	}()
	// map1 := make(map[string]string)
	var map1 map[string]string // 故意报错
	map1["李四"] = "喝酒"
}
