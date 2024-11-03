package main

//修改模版引擎的标识符
//避免恶意攻击(xss攻击)

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版(不同)
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles(".//modify_identifier/index.tmpl")
	//渲染模版
	msg := "猪猪侠"
	err = t.ExecuteTemplate(w, "index.tmpl", msg)
	if err != nil {
		fmt.Println("Execute template failed,err:", err)
		return
	}

}
func xss(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版(不同)
	//解析模版之前定义自定义函数
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles(".//modify_identifier/xss.tmpl")
	//渲染模版
	msg1 := "<script>alert(123);</script>"
	//这一点代码是恶意攻击
	msg2 := "<a href= 'http://liwenzhou.com'>liwenzhou的博客;</a>"
	err = t.Execute(w, map[string]string{
		"str1": msg1,
		"str2": msg2,
	})
	if err != nil {
		fmt.Println("Execute template failed,err:", err)
		return
	}

}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server starts failed,err:", err)
		return
	}
}
