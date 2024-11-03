package main

import "fmt"

//结构体嵌套以及结构体的字段冲突

type Address struct {
	Province   string
	City       string
	UpdateTime string
}

type Email struct {
	name       string
	UpdateTime string
}

type Person struct {
	Gender  string
	Name    string
	Age     int8
	Address //嵌套另外一个结构体 也可以使用匿名结构体
	Email
}

func main() {
	p1 := Person{
		Gender: "male",
		Name:   "Jay",
		Age:    12,
		Address: Address{
			Province:   "四川",
			City:       "南充",
			UpdateTime: "地址更新时间",
		},
		Email: Email{
			"@qq.com",
			"邮件更新时间",
		},
	}
	fmt.Printf("%#v", p1)
	// fmt.Println(p1.UpdateTime)
	//此时有两个UpdateTime,将产生冲突
	fmt.Println(p1)
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.UpdateTime)
}
