package test

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)
//没完成
//数据失真

type data struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

func TestJson(t *testing.T) {
	u1 := data{
		Name: "测试数据失真",
		Id:   math.MaxInt64,
	}
	//序列化: go语言中的数据 ——> json格式数据
	d1, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(d1))

	//反序列化:json格式数据  ——> go语言中的数据
	var u2 data
	err = json.Unmarshal(d1, &u2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u2)

}
