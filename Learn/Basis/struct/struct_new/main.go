package main

import "fmt"

//结构体构造函数：构造一个结构体实例的函数
//结构体是值类型
//当结构体较大时，构造函数一般传地址

type person struct {
	name, city string
	age        int8
}

// 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	p2 := newPerson("小王子", "北京", 15)
	fmt.Printf("%#v", p2)

}
