package main

// Gin中的中间件
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//在goroutine中只能使用c的拷贝

// handlerFunc
func handlerFunc(c *gin.Context) {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"msg": "I'm index",
	})
}

// 定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()
	c.Next() //调用后续处理函数
	//c.Abort() //阻止后续处理函数调用
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost.Nanoseconds())
	fmt.Println("m1 out...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	//c.Next() //调用后续处理函数
	//c.Abort()//阻止后续处理函数调用(只管自己)
	//c.Next和c.Abort都没有的话就是按顺序一个一个来
	c.Set("name", "小王子") //设值
	fmt.Println("m2 out...")
}
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//	或这一些其他准备工作
	return func(c *gin.Context) {
		if !doCheck {
			//存放具体的逻辑
			//是否登录的判断
			//if 是否登录用户
			//c.Next()
			//else
			c.Abort()
		} else {
			c.Next()
		}
	}
}
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name") //从上下文中取值
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{"name": name})
}

func main() {
	r := gin.Default()

	r.Use(m1, m2) //全局注册中间件函数m1,m2
	//GET is a shortcut for router.Handle("GET", path, handlers).
	r.GET("/index", handlerFunc) //先m1再handlerFunc
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Name": "shop",
		})
		c.Get("name")
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Name": "user",
		})
	})
	//为路由组注册中间件 	方法一
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}
	//为路由组注册中间件 	方法二
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	}
	r.GET("/home", indexHandler)
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return
	}
}
