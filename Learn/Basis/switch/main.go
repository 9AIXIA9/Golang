package main

import "fmt"

// switch语句特殊写法
func main() {
	//多种情况并行
	x := 10
	switch x {
	case 1, 3, 5, 7, 9:
		fmt.Println("我是奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("我是偶数")
	default:
		fmt.Println("我不知道我是什么")
	}

}
