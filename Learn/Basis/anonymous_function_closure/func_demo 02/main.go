package main

import (
	"fmt"
	"strings"
)

//闭包的扩展

// 判断是否为闭包，则需分析内部函数是否调用外部环境
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	r := makeSuffixFunc(".txt")
	x := r("莽夫")
	fmt.Println(x)

	r2 := makeSuffixFunc(".avi")
	x = r2("杨伤心")
	fmt.Println(x)

	t1, t2 := calc(100)
	fmt.Println(t1(200), t2(200))

}
