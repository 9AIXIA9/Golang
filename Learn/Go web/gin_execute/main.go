package main

//gin框架中渲染模版
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//静态文件：html页面上用到的样式文件 .css js文件 图片

func main() {
	r := gin.Default()
	//gin框架中给模版添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	//加载静态文件要在解析模版之前
	r.Static("/xxx", "./statics")
	//这里表示所有以/xxx开头的都去./statics下寻找
	//r.LoadHTMLFiles(".//templates/posts/index.tmpl",".//templates/users/index.tmpl") //解析模版
	r.LoadHTMLGlob("templates/**/*") //正则匹配:代表templates下面的目录下的所有文件

	r.GET("posts/index",
		func(ctx *gin.Context) {
			//HTML请求码
			ctx.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //渲染模版
				"title": "post的模版", //相当于map
			})
		})
	r.GET("users/index",
		func(ctx *gin.Context) {
			//HTML请求码
			ctx.HTML(http.StatusOK, "users/index.tmpl", gin.H{ //渲染模版
				"title": "<a href='https://liwenzhou.com'>李文周的博客</a>", //相当于map
			})
		})
	err := r.Run(":9090") //启动server
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
