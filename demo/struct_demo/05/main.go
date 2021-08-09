package main

import "fmt"

// 结构体继承，通过嵌套类型实现
func main() {
	var d Dog
	d.Name = "阿奇"
	d.Anm.run()
	d.work()
}

type Anm struct {
	Name string
}

func (a Anm) run() {
	fmt.Printf("%v开始跑了 \n", a.Name)
}

type Dog struct {
	age string
	Anm
}

func (d Dog) work() {
	fmt.Printf("%v在旺旺 \n", d.Name)
}
