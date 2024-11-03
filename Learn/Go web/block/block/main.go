package main

//模版继承

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模版(模版继承的方式)
	//解析模版
	t, err := template.ParseFiles("./block/block/templates/base.tmpl", "./block/block/templates/index.tmpl") //根模版在前面
	if err != nil {
		fmt.Println("parse failed,err:", err)
		return
	}
	msg := "猪猪侠"
	//渲染模版
	err = t.ExecuteTemplate(w, "index.tmpl", msg)
	if err != nil {
		fmt.Println("Execute failed,err:", err)
		return
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模版(模版继承的方式)
	//解析模版
	t, err := template.ParseFiles("./block/block/templates/base.tmpl", "./block/block/templates/home.tmpl") //根模版在前面
	if err != nil {
		fmt.Println("parse failed,err:", err)
		return
	}
	msg := "小王子"
	//渲染模版
	err = t.ExecuteTemplate(w, "home.tmpl", msg)
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
