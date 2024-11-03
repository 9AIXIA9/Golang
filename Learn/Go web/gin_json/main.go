package main

//gin框架中获得json
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json/map", func(c *gin.Context) {
		//方法一：使用map
		/*data := map[string]interface{}{
			"name":    "小王子",
			"age":     18,
			"gender":  "女",
			"message": "Hello world!",
		}*/
		data := gin.H{
			"name":    "小王子",
			"message": "诸事顺利",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/json/struct", func(c *gin.Context) {
		//方法二：结构体
		//灵活使用tag来进行定制化操作
		type msg struct {
			Name    string `json:"用户名字"`
			Age     int    `json:"age"`
			Message string `json:"message"`
			gender  string
		}
		data := msg{
			"小王子",
			18,
			"诸事顺利",
			"女",
		}
		c.JSON(http.StatusOK, data) //大写可导出，小写无法导出
	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
