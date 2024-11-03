package main

//路由 与 路由组
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问/index的GET请求会走这一处理逻辑
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	//发送数据
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	//删除
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//更新
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	//所有方法都用这个(请求方法大杂烩)
	r.Any("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "Any",
		})
		switch c.Request.Method {
		//1.
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
			//2.
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
			//...
		}
	})
	//NoRoute(用于处理未定义页面)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "We dont have this web"})
	})
	//路由组
	videoGroup := r.Group("/video") //把公用前缀提出创建路由组
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, "/video/index")
		})
		videoGroup.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, "/video/home")
		})
	}
	//路由组能进行嵌套
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
