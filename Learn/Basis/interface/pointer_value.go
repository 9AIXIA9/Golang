package main

import "fmt"

//使用值接受者实现接口和使用指针接受者实现接口的区别

type mover interface {
	move()
}

type sayer interface {
	say()
}

// 接口的嵌套
type animal interface {
	mover
	sayer
	//相当于
	//move()
	//say()
}

type person struct {
	name string
	age  int8
}

//// 使用值接受者实现接口:类型的值和类型的指针都可以保存到接口中
//func (p person) move() {
//	fmt.Printf("%s在跑\n", p.name)
//}

// 使用指针接受者实现接口
func (p *person) move() {
	fmt.Printf("%s在跑\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s在叫\n", p.name)
}

func main() {
	var m mover
	//p1 := person{ //person类型的值
	//	"小王子",
	//	18,
	//}
	p2 := &person{
		name: "小白",
		age:  16,
	}
	//m = p1 // ？ 无法赋值，因为p1是person值类型，没有实现接口mover接口
	//m.move()
	m = p2
	m.move()

	var s sayer
	s = p2
	fmt.Println(s)
	p2.say()

	var a animal
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}
