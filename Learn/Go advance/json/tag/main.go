package main

// 使用json tag指定序列化与反序列化时的行为

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name   string  `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64   `json:"age"`
	Weight float64 `json:"-"`
	Gender bool    `json:"gender,omitempty"`
	Addr   `json:"addr,omitempty"`
}
type Addr struct {
	Province string `json:"province"`
	City     string `json:"city,omitempty"`
}

func main() {
	p := Person{
		Name:   "构式",
		Age:    10,
		Weight: 38,
	}
	log.Printf("p:%#v", p)
	by, err := json.Marshal(p)
	if err != nil {
		log.Print("json marshal failed", err)
		return
	}
	log.Printf("%s", by)

	//匿名嵌套
	omitPasswordDemo()
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User
	Password *struct{} `json:"password,omitempty"`
}

type TestUser struct {
	Name     string `json:"name"`
	Password string `json:"-"`
}

// 我们需要json序列化User
// 但是不想把密码也序列化
// 又不想修改User结构体
// 这个时候我们就可以使用创建另外一个结构体PublicUser匿名嵌套原User
// 同时指定Password字段为匿名结构体指针类型
// 并添加omitempty
func omitPasswordDemo() {
	u1 := User{
		Name:     "七米",
		Password: "123456",
	}
	u2 := TestUser{
		Name:     "七米",
		Password: "123456",
	}
	b, err := json.Marshal(PublicUser{User: &u1})
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	b2, err := json.Marshal(&u2)
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}

	fmt.Printf("str:%s\n", b)  // str:{"name":"七米"}
	fmt.Printf("str:%s\n", b2) // str:{"name":"七米"}
}
