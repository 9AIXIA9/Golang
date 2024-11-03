package main

//gin框架_demo
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello golang!",
	})
}
func main() {
	r := gin.Default() //返回默认的路有引擎
	//指定用户使用GET访问/hello,	执行sayHello这个函数
	r.GET("/hello", sayHello)

	r.GET("/book", sayHello)

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//启动服务
	r.Run(":9090")
}
