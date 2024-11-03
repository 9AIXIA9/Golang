package main

//用gin获得form
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("./index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("./login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// /login post
	r.POST("/login", func(c *gin.Context) {
		//第一种方法
		//username := c.PostForm("username112")
		//password := c.PostForm("password112")
		//第二种方法
		//username := c.DefaultPostForm("username112", "somebody")
		//password := c.DefaultPostForm("password112	", "****")
		//第三种方法
		username, ok := c.GetPostForm("username112")
		if !ok {
			username = "sb"
		}
		password, ok := c.GetPostForm("password112")
		if !ok {
			password = "******"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"password": password,
		})
	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
