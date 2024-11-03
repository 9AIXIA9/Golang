package main

import "fmt"

//自定义类型和类型别名

//1.自定义类型

// MyInt 基于int的自定义类型
type MyInt int

//2.类型别名

// NewInt int类型别名
type NewInt = int

func main() {
	var x MyInt
	fmt.Printf("type:%T Value:%v\n", x, x)
	var y NewInt
	fmt.Printf("type:%T Value:%v\n", y, y)
}
