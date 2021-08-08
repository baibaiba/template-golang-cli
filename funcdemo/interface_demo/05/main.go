package main

import "fmt"

// 一个结构体实现多个接口
func main() {
	var d = Dog{}
	var d1 Ainterface = d
	var d2 Binterface = d
	d1.start()
	d2.stop()
}

type Ainterface interface {
	start()
}

type Binterface interface {
	stop()
}

type Dog struct {
}

func (d Dog) start() {
	fmt.Println("start")
}

func (d Dog) stop() {
	fmt.Println("stop")
}
