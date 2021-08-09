package main

import (
	"fmt"
)

// 单向管道
func main() {
	// 1.在默认情况下，管道是双向的
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	fmt.Println(<-ch1, <-ch1, <-ch1)

	// 2.声明管道为只写
	ch2 := make(chan<- int, 3)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	// <-ch2 // 报错（invalid operation: <-ch2 (receive from send-only type chan<- int)）

	// 2.声明管道为只读
	// ch3 := make(<-chan int, 3)
	// ch3 <- 1 // 报错（invalid operation: ch3 <- 1 (send to receive-only type <-chan int)）
}
