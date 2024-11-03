package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn/controller"
)

func OpenWeb() error {
	var engine = gin.Default()
	User := engine.Group("/user")
	//加载templates中所有模板文件, 使用不同目录下名称相同的模板,注意:一定要放在配置路由之前才得行
	engine.LoadHTMLGlob("html/*")
	User.POST("/login", controller.LoginUser)   //用户浏览器直接进的登录页面
	User.POST("/enroll", controller.EnrollUser) //用户浏览器直接进的注册页面
	User.GET("/login", controller.LoginUser)    //用户在提交登录表单过后，对此反应的方法
	User.GET("/enroll", controller.EnrollUser)  //用户提交注册表单过后，对此进行反应的方法
	err := engine.Run(":9090")
	if err != nil {
		fmt.Println("server run failed,err:", err)
		return err
	}
	return nil
}
