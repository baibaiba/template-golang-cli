package main

import (
	"fmt"
)

// channel
func main() {
	/**
	channel的概念：管道时golang在语言级别上提供的goroutine间的通讯方式，我们可以使用channel在多个goroutine之间传递消息。如果说goroutine是go程序并发的执行体，channel就是他们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
	golang的并发模型是csp(communication sequential processes)，提倡【通过通信共享内存】，而不是【通过共享内存实现通信】
	go语言中的管道是一种特殊的类型。管道像一个传送带或者队列，总是遵循【先入先出（first in first out）】的原则，保证收发数据的顺序。
	每一个管道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型
	*/

	/**
	1.channel的类型
	channel是一种类型，一种引用类型。声明管道类型的格式如下
	var 变量 chan 元素类型
	*/
	// var ch chan int

	/**
	2.创建channel
	声明管道后需要使用make()函数初始化之后才能使用
	*/
	// var ch1 = make(chan int, 3)

	/**
	3.channel的操作
	管道有发送（send）、接收（receive）、关闭（close）三种操作
	*/
	var ch2 = make(chan int, 3)
	ch2 <- 10 // 发送数据
	ch2 <- 20
	ch2 <- 30
	a := <-ch2 // 接收数据并赋值
	<-ch2      // 接收数据不赋值
	fmt.Println(a)
	fmt.Printf("channel值：%v，容量：%v，长度：%v \n", ch2, cap(ch2), len(ch2))

	fmt.Println("--------------------------------")
	/**
	4.channel的类型（引用数据类型）
	*/
	var ch3 = make(chan int, 3)
	ch3 <- 10
	ch4 := ch3
	ch4 <- 20
	a1 := <-ch3
	a2 := <-ch3
	fmt.Println(a1, a2)

	/**
	5.channel的阻塞
	*/
	// var ch5 = make(chan int, 1)
	// ch5 <- 88
	// ch5 <- 89 // 报错，没有容量可以存储（fatal error: all goroutines are asleep - deadlock!）

	// var ch6 = make(chan int, 2)
	// ch6 <- 88
	// ch6 <- 89
	// a3 := <-ch6
	// a4 := <-ch6
	// a5 := <-ch6 // 没有足够的容量可以取出（fatal error: all goroutines are asleep - deadlock!）
	// fmt.Println(a3, a4, a5)

	fmt.Println("--------------------------------")
	/**
	6.channel的遍历
	*/
	// var ch7 = make(chan int, 2)
	// ch7 <- 88
	// ch7 <- 89
	// for v := range ch7 { // 管道没有下标
	// 	fmt.Println(v)
	// }
	// 光执行上面的代码会报错（fatal error: all goroutines are asleep - deadlock!）
	// 注意：使用range遍历前channel必须关闭channel
	// 正确写法
	var ch7 = make(chan int, 2)
	ch7 <- 88
	ch7 <- 89
	close(ch7)           // 我的理解：关闭管道的入口方向，不让再写入数据
	for v := range ch7 { // 管道没有下标
		fmt.Println(v)
	}

	// 使用普通for遍历channel可以关闭也可以不关闭channel
	var ch8 = make(chan int, 2)
	ch8 <- 999
	ch8 <- 998
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch8)
	}
}
