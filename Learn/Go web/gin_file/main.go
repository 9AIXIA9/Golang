package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)

	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件

		f, err := c.FormFile("f1") //从请求中获取携带的参数
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			//将读取的文件保存在本地(服务端)
			//dst := fmt.Sprintf(".%s", f.Filename)
			dst := path.Join("./", f.Filename)
			_ = c.SaveUploadedFile(f, dst) //第二个参数就是文件保存位置
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})

		}

	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("server run failed,err", err)
	}
}
