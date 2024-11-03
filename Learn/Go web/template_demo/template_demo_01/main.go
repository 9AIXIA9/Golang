package main

//template_demo_01
import (
	"fmt"
	"html/template"
	"net/http"
)

// 遇事不决，先写注释
func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	t, err := template.ParseFiles("./template_demo_01/hello.tmpl")
	//不要刻舟求剑
	// 此处./只代表项目的本目录
	if err != nil {
		fmt.Println("Parse template failed,err:", err)
		return
	}
	//3.渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Execute failed,err:", err)
		return
	}
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed,err:", err)
		return
	}
}
