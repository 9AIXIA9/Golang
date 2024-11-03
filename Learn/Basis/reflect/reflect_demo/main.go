package main

import (
	"fmt"
	"reflect"
)

// reflect demo
func reflectType(x interface{}) {
	//我是不知道别人调用我这个函数的时候会传进来什么类型
	//方法1：类型断言
	//方法2：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj)
}
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() //拿到值对应的传入类型种类
	fmt.Println(k)
	fmt.Printf("%v,%T\n", v, v)
	//如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		//把反射取到的值转换为float32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Int32:
		//把反射取到的值转换为Int32类型的变量
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)
	default:
		panic("unhandled default case")
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	//Elem()根据指针取对应的值
	k := v.Elem().Kind() //拿到值对应的传入类型种类
	switch k {
	case reflect.Float32:
		//把反射取到的值转换为float32类型的变量
		v.Elem().SetFloat(1.234)
	case reflect.Int32:
		v.Elem().SetInt(100)
	default:
		panic("unhandled default case")
	}
}

type cat struct {
}

type dog struct {
}

func main() {
	//var a float32 = 1.234
	////reflectType(a)
	//reflectValue(a)
	//var b int32 = 8
	////reflectType(b)
	//reflectValue(b)
	////结构体类型
	//var c cat
	//reflectType(c)
	//var d dog
	//reflectType(d)
	////slice没有name
	//var e []int
	//reflectType(e)
	//var f []string
	//reflectType(f)
	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)
}
