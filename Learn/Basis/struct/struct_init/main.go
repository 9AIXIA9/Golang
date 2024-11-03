package main

import "fmt"

//结构体的初始化

type person struct {
	name, city string
	age        int8
}

func main() {
	//1.键值对初始化
	p3 := person{
		name: "小王子",
		city: "北京",
		age:  12,
	}
	p4 := &person{
		name: "小王子",
		city: "北京",
		age:  12,
	}
	fmt.Printf("%#v", p3)
	fmt.Printf("%#v", p4)
	//2.值的列表初始化
	p5 := person{
		"小王子",
		"北京",
		12,
	}
	fmt.Printf("%#v", p5)
	fmt.Printf("%+v\n", p5)
	//以名字加数据（超好用）
}
