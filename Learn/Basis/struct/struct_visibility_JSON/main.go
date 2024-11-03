package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性和JSON的序列化

// 一个go语音包中定义的标识符首字母是大写，则是对外部可见的
// 如果一个结构体的字段的首字母是大写的，那么该字段对外可见
type student struct {
	ID   int
	Name string
}

type class struct {
	Title    string    `json:"title" db:"title"` //结构体的tag,改变json中的名称，改变db数据库里面的名称
	Students []student `json:"Students" db:"student"`
}

// student的构造函数
func newStudent(ID int, Name string) student {
	return student{
		ID:   ID,
		Name: Name,
	}
}
func main() {
	//创建一个班级变量c1
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}
	//往班级中添加学生
	for i := 0; i < 10; i++ {
		//造十个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	//fmt.Printf("%#v\n", c1)
	//JSON的序列化:go语言中的数据转变为 -> JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	//fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	//JSON的反序列化: JSON格式的字符串 -> go语言中的数据转变为
	var c2 class
	err = json.Unmarshal([]byte(data), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed, err:", err)
		return
	}
	//fmt.Printf("%#v", c2)
	fmt.Println("反序列成功")
}
