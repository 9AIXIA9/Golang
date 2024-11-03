package main

import "fmt"

//变量和常量

// 常量声明
const pai = 3.1415926
const (
	birthday = 617
	birth
	theBirthday
	//当不赋初值时，默认跟前一常量一致
)

/*
iota用于记录常量声明次数(枚举)
const出现则iota重新归为领
每增加一行变量则+1
一行！
*/
const (
	n1 = iota
	n2
	n3
	_
	//匿名变量依旧占据
	n4 = 100
	n5 = iota
)

// 全局变量
var x = 10
var y = "little prince"

// 变量
func main() {
	fmt.Println("圆周率是", pai)
	fmt.Println(birthday, "birthday is", theBirthday, birth)
	fmt.Println(n1, n2, n3, n4)
	//输出元素之间有空格
	fmt.Println(x, y, y, y)
	//标准变量声明
	var name string
	var age int
	var judgement bool
	fmt.Println(name, age, judgement)
	//批量变量声明
	var (
		a int
		b string
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	//声明变量同时指定初始值
	var (
		e int8 = 6
		f      = "我真帅!"
	)
	//类型推导
	var g, k = "彭于晏", 10
	fmt.Println(e, f, g, k)
	//短变量声明
	m := 10
	fmt.Println(m)
	/*匿名变量
	用_下划线占据变量位置
	*/

}
