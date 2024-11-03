package main

import "fmt"

// 指针
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
func main() {
	x := 1
	modify1(x)
	fmt.Println(x)
	modify2(&x)
	fmt.Println(x)
	//指针必须初始化
	////用new初始化a
	//var a *int
	//a = new(int)

}
