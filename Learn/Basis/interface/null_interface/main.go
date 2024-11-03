package main

import "fmt"

//空接口
//接口中没有定义任何需要实现的方法时，该接口就是空接口
//任意类型实现了空接口 ——> 空接口变量可以储存任意值

//空接口一般不需要提前定义

type xxx interface {
}

//空接口的应用：
//1.空接口类型作为函数的参数
//2.空接口作为map的value

func main() {
	//1.
	//定义一个空接口变量
	var x xxx
	x = "false"
	fmt.Println(x)
	x = 23
	fmt.Println(x)
	x = false
	fmt.Println(x)
	//2.
	var m = make(map[string]interface{}, 16)
	m["name"] = "娜扎"
	m["age"] = 19
	m["hobby"] = []string{"篮球", "足球", "乒乓球"}
	fmt.Println(m)
	//类型断言	猜的类型不对时会返回一个布尔值，ok = false ，ret = 猜x的类型的零值
	ret := x.(bool)
	fmt.Println(ret)
	ret1, ok := x.(string)
	if ok {
		fmt.Println("x是string类型", ret1)
	} else {
		fmt.Println("x不是string类型")
	}
	//使用Switch语句进行类型断言
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型，value:%v\n", v)
		break
	case bool:
		fmt.Printf("是布尔类型，value:%v\n", v)
	case int:
		fmt.Printf("是整型，value:%v\n", v)
	default:
		fmt.Printf("猜不到了，value:%v\n", v)
	}
}
