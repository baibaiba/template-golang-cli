package main

import "fmt"

// 在golang中没有类的概念，但是可以给类型（结构体、自定义类型）定义方法
// 所谓方法就是定义了接收者的函数。（就像在Java类里面定义方法一样）
// 接收者的概念就类似与其它语言中的this或者self
func main() {
	p1 := Person{
		"张三",
		10,
		"男",
	}

	p1.printInfo()
	p1.setInfo("李四", 20)
	p1.printInfo()
}

type Person struct {
	Name string
	Age  int
	Sex  string
}

/**
方法的定义格式如下：
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	函数体
}
*/
func (p Person) printInfo() {
	fmt.Printf("姓名：%v,年龄：%d \n", p.Name, p.Age)
}

/**
 接收者变量：接收者中的参数变量在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名，
			 例如：Person类型的接收者变量的命名为p
 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型
 方法名、参数列表、返回参数：具体格式与函数相同
*/

// 注意:如下所示，调用方法的对象的值是不会改变的；因为结构体是值类型，会先赋值给p,修改的也是p
// 如果要修改调用对象的值，那么接收者类型必须是引用类型
func (p *Person) setInfo(name string, age int) {
	p.Name = name
	p.Age = age
}
