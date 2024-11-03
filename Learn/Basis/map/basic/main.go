package main

import (
	"fmt"
)

// map(映射)
func main() {
	//光声明不初始化,但是不初始化,则为nil
	var a map[string]int
	if a == nil {
		fmt.Println("a是一个nil")
	}
	//make的初始化
	a = make(map[string]int, 8)
	//map中添加键值对
	a["彭于晏"] = 100
	a["李白"] = 200
	fmt.Println(a)
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("%T\n", a)
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("%T\n", b)
	//判断某个键是否存在
	value, ok := a["杜甫"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("杜甫在a中存在")
	} else {
		fmt.Println("杜甫不在a中")
	}
	value, ok = a["李白"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("李白在a中存在")
	} else {
		fmt.Println("李白不在a中")
	}
	//map的遍历
	//map是无序的，键值对与添加顺序无关
	for v, i := range a {
		fmt.Println(v, i)
	}
	//只遍历数值
	for v := range a {
		fmt.Println(v)
	}
	//删除键值对
	a["我是帅哥"] = 300
	fmt.Println(a)
	delete(a, "我是帅哥")
	fmt.Println(a)
}
