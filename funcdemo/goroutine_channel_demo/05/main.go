package main

import "fmt"

// 在有些时候我们需用从多个管道中接收数据。这个时候就可以用的golang中给我们提供的select多路复用
func main() {

	ch1 := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		ch1 <- i
	}

	ch2 := make(chan string, 5)
	for i := 1; i <= 5; i++ {
		ch2 <- fmt.Sprintf("hello %d", i)
	}

	// 传统的方法获取两个channel的数据只有顺序执行
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-ch1)
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-ch2)
	// }

	// 使用多路复用，语法如下：
	for {
		select {
		case v := <-ch1: // 从ch1接收数据，如果没有数据可接受执行别的case
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		default: // 如果所有的channel数据接收完成，执行
			fmt.Println("数据读取完毕")
			return
		}
	}
	// 注意，使用多路复用不需要关闭channel,
}
