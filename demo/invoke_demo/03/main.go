package main

import (
	"fmt"
	"reflect"
)

// 结构体反射获取值，获取方法，修改值，调用方法
func main() {
	p1 := Person{"李四", 20}
	// PrintField(p1)
	// PrintMethod(&p1) // 修改必须传指针类型
	SetFieldValue(&p1, "赵六")
	fmt.Printf("%#v \n",p1)
}

func PrintField(x interface{}) {
	// 类型变量
	t := reflect.TypeOf(x)
	// 值变量
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}

	// 第一种获取属性的方式
	f := t.Field(0) // 指定顺序，就是从类定义的字段从上到下获取
	fmt.Printf("%#v \n", f)
	fmt.Printf("属性名称:%v 属性类型:%v 属性Tag:%v 另外一个属性Tag:%v \n", f.Name, f.Type, f.Tag.Get("json"), f.Tag.Get("form")) // 注意tag的获取方式

	// 第二种根据属性名获取
	f1, ok := t.FieldByName("Age")
	if ok {
		fmt.Printf("属性名称:%v 属性类型:%v 属性Tag:%v  \n", f1.Name, f1.Type, f1.Tag.Get("json")) // 注意tag的获取方式
	}

	// 获取结构体中有几个属性
	num := t.NumField()
	fmt.Printf("结构体中有%v个属性", num)

	// 第一种获取属性值
	f2, ok := t.FieldByName("Name")
	if ok {
		fmt.Printf("属性名称:%v 属性类型:%v 属性值:%v  \n", f2.Name, f2.Type, v.FieldByName("Name")) // 注意属性值需要使用变量类型
	}

	fmt.Println("---------------------------------------")
	// 第二种获取属性值
	for i := 0; i < num; i++ {
		f3 := t.Field(0)
		fmt.Printf("属性名称:%v 属性类型:%v 属性值:%v  \n", f3.Name, f3.Type, v.Field(i)) // 注意属性值需要使用变量类型
	}
}

func PrintMethod(x interface{}) {
	// 类型变量
	t := reflect.TypeOf(x)
	// 值变量
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}

	// 获取方法的第一种方式
	m := t.Method(0) // 这里获取顺序是根据ascll排序
	fmt.Printf("%#v \n", m)
	fmt.Printf("方法名称：%v 方法类型：%v \n", m.Name, m.Type)

	// 获取方法的第二种方式
	m1, ok := t.MethodByName("PrintInfo")
	if ok {
		fmt.Printf("方法名称：%v 方法类型：%v \n", m1.Name, m1.Type)
	}

	// 获取类里面的方法数量
	fmt.Println("结构体内的方法数量", t.NumMethod())

	// 调用结构体中的方法(注意使用值类型执行方法)
	info := v.MethodByName("GetInfo").Call(nil) // 传nil代表调用时不穿任何参数
	fmt.Println(info)

	fmt.Printf("修改前%#v \n", x)
	var params []reflect.Value
	params = append(params, reflect.ValueOf("王五"), reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params)
	fmt.Printf("修改后%#v \n", x)
}

func SetFieldValue(x interface{}, newName string) {
	// 类型变量
	t := reflect.TypeOf(x)
	// 值变量
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体")
		return
	}

	name := v.Elem().FieldByName("Name")
	name.SetString(newName)
	fmt.Println("修改name值为", newName)
}

type Person struct {
	Name string `json:"name" form:"username"`
	Age  int    `json:"age"`
}

func (p Person) PrintInfo() {
	fmt.Printf("%#v \n", p)
}

func (p Person) GetInfo() string {
	return p.Name
}

func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}
