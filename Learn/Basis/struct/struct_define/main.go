package main

import "fmt"

//结构体

type person struct {
	name, city string
	age        int8
}

func main() {
	var p person
	p.name = "柠檬"
	p.city = "四川"
	p.age = 23
	fmt.Printf("p = %#v\n", p)
	//匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.married = false
	user.name = "小新"
	fmt.Printf("user : %#v\n", user)

}
