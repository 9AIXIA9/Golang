package main

import "fmt"

//任何类型都可以添加方法，但是只能为本地类型添加方法

type myInt int

func (m myInt) sayHi() {
	fmt.Println("Hi!")
}

func main() {
	var m1 myInt
	m1 = 100
	m1.sayHi()
}
