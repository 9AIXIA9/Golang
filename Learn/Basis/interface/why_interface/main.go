package main

import "fmt"

//为什么需要接口

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

type dog struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

// 接口不管你是什么类型，只管具体实现方法

// 定义一个类型，只要定义了say这个方法，都可以算作sayer类型
type sayer interface {
	say()
}

// 打的函数
func da(arg sayer) {
	//不管传进来什么都要打，打ta就叫，叫就要执行方法
	arg.say()
}

func main() {
	c1 := cat{}
	d2 := dog{}
	da(c1)
	da(d2)
	//没有da函数就没法当做sayer
	//p3 := person{
	//	name string
	//}
	//可以用接口变量存储所有实现了接口类型的变量
	c1.name = "小四"
	var s4 sayer = c1
	fmt.Println(s4)
	fmt.Println(c1)
}
