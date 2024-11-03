package main

import "fmt"

//方法的定义实例

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 是一个Person类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// 定义方法
// func (接收者)函数名(){}

// Dream 是为Person类型定义的方法
func (p *Person) Dream() {
	fmt.Printf("%v的梦想是在%v岁时征服星辰大海\n", p.name, p.age)
}

// SetAge 修改Person年龄的方法
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// 指针接收者指的是接收者的类型是指针

// SetAge2 是一个使用值来当接受者的类型

func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("路飞", 18)
	(*p1).Dream()
	p1.Dream()
	p1.SetAge(20)
	p1.Dream()
	p1.SetAge2(30)
	p1.Dream()
	//传入的是值类型则进行值拷贝
	//1.要修改接受者的值时
	//2.接收者是拷贝类型比较大的对象时
	//3.保证一致性，若有其他方法使用指针接受者
	//这三种情况用指针接受者
}
