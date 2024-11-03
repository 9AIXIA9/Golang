package main

import (
	"fmt"
	"reflect"
)

//结构体反射

type student struct {
	Name  string `json:"name" ini:"S_Name"`
	Score int    `json:"score" ini:"S_Score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 12,
	}
	//通过反射去获取结构体的所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("Name:%v kind:%v\n", t.Name(), t.Kind())
	//遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		//根据结构体的索引去取字段
		fileObj := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v\n", fileObj.Name, fileObj.Type, fileObj.Tag)
		fmt.Println(fileObj.Tag.Get("json"), fileObj.Tag.Get("ini"))
	}
	//根据名字去取结构体中的字段
	filed, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v\n", filed.Name, filed.Type, filed.Tag)
	}
	//方法个数
	fmt.Println(t.NumMethod())
}
