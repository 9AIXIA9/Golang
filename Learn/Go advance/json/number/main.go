package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

//优雅处理字符串格式的数字
//整数变浮点数

type Card struct {
	ID    int64   `json:"id,string"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}

func main() {
	intAndStringDemo()
	useNumberDemo()
}

func intAndStringDemo() {
	jsonStr1 := `{"id": "1234567","score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) // c1:main.Card{ID:1234567, Score:88.5}
}

// useNumberDemo 使用json.UseNumber
// 解决将JSON数据反序列化成map[string]interface{}时
// 数字变为科学计数法表示的浮点数问题
// 因为在 JSON 协议中是没有整型和浮点型之分的,它们统称为number
// json字符串中的数字经过Go语言中的json包反序列化之后都会成为float64类型
func useNumberDemo() {
	type student struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	s := student{ID: 123456789, Name: "q1mi"}
	b, err := json.Marshal(s)
	if err != nil {
		log.Print(err)
		return
	}
	var m map[string]interface{}
	// decode
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Printf("id:%#v\n", m["id"])     // 1.23456789e+08
	fmt.Printf("id type:%T\n", m["id"]) //float64

	// use Number decode
	reader := bytes.NewReader(b) // 把 json 数据转化为 io reader 类型
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()      //告诉解码器在解码 JSON 数字时使用 json.Number 类型，而不是默认的 float64
	err = decoder.Decode(&m) //将json数据解码到 m 中
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Printf("id:%#v\n", m["id"]) // "123456789"
	//json number 是一个字符串，保留了原始的数字表示
	//json.Number().Int64() json.Number().Float64() 此类型是可以调用方法的
	fmt.Printf("id type:%T\n", m["id"]) // json.Number
}
