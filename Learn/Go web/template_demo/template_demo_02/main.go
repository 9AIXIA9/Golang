package main

//template_demo_02

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name,
	Gender string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./template_demo/template_demo_02/hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	//渲染模版
	u1 := User{ //u1.Name
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"Name":   "小小王子",
		"Age":    28,
		"Gender": "女",
	}
	hobbyList := []string{"羽毛球", "游戏", "看番"}
	err = t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("Execute failed,err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server starts failed,err:", err)
		return
	}
}
