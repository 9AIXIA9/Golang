package main

import "fmt"

//结构体指针

type person struct {
	name, city string
	age        int8
}

func main() {
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	//1.第一种写法
	//(*p2).name = "小王子"
	//(*p2).city = "北京"
	//(*p2).age = 11
	//第二种写法（go语言不存在指针操作）
	p2.name = "小王子"
	p2.city = "北京"
	p2.age = 11
	fmt.Printf("%#v", p2)

}
