package main

import "fmt"

// golang中的interface
func main() {
	// 定义：golang中的接口是一种抽象数据类型，golang中接口定义了对象的行为规范，中定义规范不实现。
	// 接口中定义的规范由具体的对象实现

	// --------------------------------------------------
	// 对象实现接口的方式
	p1 := Phone{
		Name: "华为",
	}
	p1.func1()
	p1.start("")
	p1.stop(0)
	fmt.Printf("%#v,类型是：%T \n", p1, p1)
	// 上面的这种方式，虽然结构体Phone类中有定义和Usber接口类中一样的方法，但是不是Phone实现了Usber
	p2 := Phone{
		Name: "小米",
	}
	var p3 Usber = p2
	// p3.func1() 这个时候就不能再访问Phone中的func1方法了
	p3.start("")
	p3.stop(0)
	fmt.Printf("%#v,类型是：%T \n", p3, p3)
	// 上面的方式就是Phone结构体实现了Usber接口，也可以通过下面的方式简写
	var p4 Usber = Phone{
		Name: "vivo",
	}
	p4.start("")
	p4.stop(0)
	fmt.Printf("%#v,类型是：%T \n", p4, p4)
	// 注意：实现接口必须实现接口中的所以方法，否则会报错
}

// 自定义一个接口类型(接口类名后建议加er后缀，表示接口类型)
type Usber interface {
	start(string) string
	stop(int) int
}

// 如果接口里面有方法，必须通过结构体或者自定义类型实现

// Phone结构体实现Usber接口
type Phone struct {
	Name string
}

// Phone结构体内的自定义方法
func (p Phone) func1() {
	fmt.Println("main Phone func1")
}

// Phone结构体实现Usber接口方法
func (p Phone) start(x string) string {
	fmt.Println("main Phone start")
	return x
}

func (p Phone) stop(x int) int {
	fmt.Println("main Phone stop")
	return x
}

// Camera结构体实现Usber接口
type Camera struct {
}
