package main

import "fmt"

//结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%v会动\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //匿名嵌套，且嵌套类型是结构体指针
}

func (a *Dog) wang() {
	fmt.Printf("%v会汪汪叫\n", a.name)
}

func main() {
	d1 := Dog{
		Feet:   4,
		Animal: &Animal{name: "旺财"},
	}
	d1.move()
	d1.wang()
}
