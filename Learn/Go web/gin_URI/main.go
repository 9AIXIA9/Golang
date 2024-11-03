package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取请求中的path(URI)参数
func main() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(c *gin.Context) {
		//  第一个/赋值给了name,第二个/赋值给了age
		//ex：http://127.0.0.1:9090/小王子/12 -> name = 小王子   age = 12

		//获取路径参数
		name := c.Param("name")
		age := c.Param("age") //返回的都是字符串类型
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})

	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
