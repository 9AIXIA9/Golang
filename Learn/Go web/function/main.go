package main

//template自定义函数
import (
	"fmt"
	"html/template"
	"net/http"
)

// 定义一个函数praise(自定义函数要么只有一个返回值,要么第二个返回值必须是err)
func praise(name string) (string, error) {
	return name + "帅死了", nil
}

func f1(w http.ResponseWriter, r *http.Request) {

	//定义模版

	t := template.New("f.tmpl") //创建一个名称为f.tmpl的模版对象,name要与模版名字对应
	//告诉模版引擎,我多了一个praise函数
	t.Funcs(template.FuncMap{
		"kua": praise,
	})
	//解析模版

	_, err := t.ParseFiles(".//function/f.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	name := "小王子"
	//渲染模版
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("execute failed,err:", err)
		return
	}
}
func demo1(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles(".//function/t.tmpl", ".//function/ul.tmpl") //大的在前，小的在后(被包含的在后)
	if err != nil {
		fmt.Println("parse failed,err:", err)
		return
	}

	//渲染模版
	name := "新之助"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Execute failed,err:", err)
		return
	}
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server starts failed,err:", err)
		return
	}
}
