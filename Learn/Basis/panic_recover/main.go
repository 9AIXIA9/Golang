package main

import "fmt"

// panic和recover函数的使用
func a() {
	fmt.Println("func a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}
func main() {
	a()
	b()
	c()
}
