package main

//未使用模版继承

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles(".//block/without_block/index.tmpl")
	if err != nil {
		fmt.Println("parse failed,err:", err)
		return
	}
	msg := "猪猪侠"
	//渲染模版
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Execute failed,err:", err)
		return
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles(".//block/without_block/home.tmpl")
	if err != nil {
		fmt.Println("parse failed,err:", err)
		return
	}
	msg := "小王子"
	//渲染模版
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Execute failed,err:", err)
		return
	}

}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server starts failed,err:", err)
		return
	}
}
