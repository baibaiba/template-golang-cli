package main

import (
	"fmt"
	"reflect"
)

//  reflect.TypeOf()获取任意变量的类型
func main() {
	a := 10
	b := 1.1
	c := "golang"
	d := true
	typeOf(a)
	typeOf(b)
	typeOf(c)
	typeOf(d)

	e := Person{"李四"}
	f := []int{123, 234, 5345}
	g := [3]string{"A", "B", "C"}
	typeOf(e)
	typeOf(f)
	typeOf(g)

	h := make(map[string]string)
	typeOf(h)

	i := &Person{"李四"}
	typeOf(i)
}

// 部分类型是没有类型名称的
func typeOf(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("类型：%v,类型名称：%v,类型种类：%v \n", v, v.Name(), v.Kind())
}

type Person struct {
	Name string
}
