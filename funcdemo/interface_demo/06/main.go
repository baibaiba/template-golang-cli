package main

import "fmt"

// 接口嵌套
func main() {
	var d = Dog{}
	var d1 Animaler = d
	d1.start()
	d1.stop()
}

type Ainterface interface {
	start()
}

type Binterface interface {
	stop()
}

// 接口嵌套
type Animaler interface {
	Ainterface
	Binterface
}

type Dog struct {
}

func (d Dog) start() {
	fmt.Println("start")
}

func (d Dog) stop() {
	fmt.Println("stop")
}
