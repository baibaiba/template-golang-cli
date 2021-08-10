package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 1
	b := "str"
	c := false
	d := 1.23
	reflectValueOf(a)
	reflectValueOf(b)
	reflectValueOf(c)
	reflectValueOf(d)

	fmt.Println("------------------------------")
	e := 1
	SetValue(&e) // 使用反射修改值必须使用指针类型
	fmt.Println(e)
	f := false
	SetValue(&f)
	fmt.Println(f)
}

// 获取原始值
func reflectValueOf(x interface{}) {
	// 返回一个reflect.Value方法
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Int:
		fmt.Printf("int类型,原始值：%v \n", v.Int())
	case reflect.String:
		fmt.Printf("String类型,原始值：%v \n", v.String())
	case reflect.Bool:
		fmt.Printf("Bool类型,原始值：%v \n", v.Bool())
	case reflect.Float32:
		fmt.Printf("Float32类型,原始值：%v \n", v.Float())
	case reflect.Float64:
		fmt.Printf("Float32类型,原始值：%v \n", v.Float())
	default:
		fmt.Println("未知类型")
	}
}

// 修改变量值
func SetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// v.Elem().Kind() 获取指针的原始类型
	switch v.Elem().Kind() {
	case reflect.Bool:
		v.Elem().SetBool(true)
	case reflect.Int:
		v.Elem().SetInt(120)
	case reflect.String:
		v.Elem().SetString("golang")
	default:
		fmt.Println("未知类型")
	}
}
