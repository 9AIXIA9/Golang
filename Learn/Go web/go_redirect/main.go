package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "OK",
		//})
		c.Redirect(http.StatusMovedPermanently, "https://www.sogou.com/")
		//跳到另外一个网站
	})
	//路由重定向
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //把请求的URI修改
		r.HandleContext(c)        //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	_ = r.Run(":8080")
}
