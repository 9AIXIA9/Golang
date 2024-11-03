package main

//querystring

//GET请求 UR?后面是querystring的参数
//key=value格式，多个key-value使用&连接
//eq：http://127.0.0.1:9090/web?query=sb&age=18
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器发送的请求携带的query string
		//name := c.Query("query值为") //通过Query函数获取请求中携带的querystring参数，在/web?query值为=中传入参数
		//name := c.DefaultQuery("query", "somebody") //找不到就用somebody
		name, ok := c.GetQuery("query") //取到返回(值,true)，取不到第二个参数就返回false
		if !ok {
			//取不到
			name = "somebody"
		}
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": "你找的是" + name,
			"age":  "年龄是" + age,
		})
	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("server run failed,err：", err)
		return
	}
}
