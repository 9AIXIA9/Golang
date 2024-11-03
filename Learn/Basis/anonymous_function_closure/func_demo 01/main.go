package main

import (
	"fmt"
)

// 匿名函数和闭包
// 定义一个返回函数的函数
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
		//内部找不到变量就从外部找
	}
}

func main() {
	////匿名函数
	//func() {
	//	fmt.Println("匿名函数")
	//}()
	r := a("ncu")
	//r此时就是一个闭包
	//闭包 = 函数 + 环境
	r() //相当于执行了a函数内部的函数

}
