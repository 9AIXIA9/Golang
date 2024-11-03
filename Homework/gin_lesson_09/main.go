package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	egn := gin.Default()
	//定义模版
	//加载静态文件要在解析模版之前
	egn.Static("/xxx", "../statics/template-mundana-bootstrap-html-master/assets")
	//这里表示所有以/xxx开头的都去./statics下寻找
	//解析模版
	egn.LoadHTMLGlob("templates/*")
	egn.GET("/article",
		func(ctx *gin.Context) {
			//HTML请求码
			ctx.HTML(http.StatusOK, "article.html", nil)
		})

	egn.GET("/about",
		func(ctx *gin.Context) {
			//HTML请求码
			ctx.HTML(http.StatusOK, "about.html", nil)
		})
	egn.GET("/category",
		func(ctx *gin.Context) {
			//HTML请求码
			ctx.HTML(http.StatusOK, "category.html", nil)
		})
	egn.GET("/docs",
		func(ctx *gin.Context) {
			//HTML请求码
			ctx.HTML(http.StatusOK, "docs.html", nil)
		})
	egn.GET("/index",
		func(ctx *gin.Context) {
			//HTML请求码
			ctx.HTML(http.StatusOK, "index.html", nil)
		})
	//渲染模版
	err := egn.Run(":9090")
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
