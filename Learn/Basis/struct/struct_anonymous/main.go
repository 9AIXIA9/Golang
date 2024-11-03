package main

import "fmt"

//结构体的匿名字段

type Person struct {
	string
	int8
}

// 同一类型不能重复
func main() {
	p1 := Person{
		"小王子",
		56,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int8)
}
