package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin绑定(bind)

type UserInfo struct {
	Name     string `form:"username" json:"username"` //需要大写,tag字段要一一对应
	Password string `form:"password" json:"password	"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(c *gin.Context) {
		//原始方法
		//name := c.Query("username")
		//password := c.Query("password")
		//u := userInfo{
		//	name:     name,
		//	password: password,
		//}
		var u UserInfo          //声明变量
		err := c.ShouldBind(&u) //值传递,修改数据需传指针
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo          //声明变量
		err := c.ShouldBind(&u) //值传递,修改数据需传指针
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.POST("/json", func(c *gin.Context) {
		var u UserInfo          //声明变量
		err := c.ShouldBind(&u) //值传递,修改数据需传指针
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("server failed,err:", err)
		return
	}
}
