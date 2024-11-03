package api

import (
	"life/internal/controllers"
	"life/pkg/middlewares/auth"
	"life/pkg/middlewares/ratelimit"
	"time"

	"golang.org/x/time/rate"
)

const (
	signupRate     = 5 * time.Second
	signupNum      = 1
	signupCleanDur = time.Minute
	loginRate      = 10 * time.Second
	loginNum       = 3
	loginCleanDur  = time.Minute
	limitTokenDur  = time.Second
	limitTokenNum  = 1
)

// v1系的路由
func setUpV1() {
	// 创建不同的限制器
	signupLimiter := ratelimit.NewIPRateLimiter(rate.Every(signupRate), signupNum, signupCleanDur)
	loginLimiter := ratelimit.NewIPRateLimiter(rate.Every(loginRate), loginNum, loginCleanDur)

	v1 := r.Group("api/v1")

	v1.POST("/signup", ratelimit.IPRLMiddleware(signupLimiter), controllers.SignUpHandler)
	v1.POST("/login", ratelimit.IPRLMiddleware(loginLimiter), controllers.LoginHandler)

	//限制登录用户才可访问
	//限制访问人数
	v1.Use(auth.JWTMiddleware(), ratelimit.TokenRLMiddleware(limitTokenDur, limitTokenNum))
	{
		//登录后可访问的路由
		//更新个人信息
		v1.POST("/update/information", controllers.UpdateInfoHandler)
		//v1.POST("/update/password", controllers.UpdateInfoHandler) //更新个人密码
		//查询能力面板
		v1.GET("/capability", controllers.GetCapabilityHandler)
		//建立个人能力面板
		v1.POST("/capability/insert", controllers.InsertCapabilityHandler)
		//改变能力面板
		v1.POST("/capability/change", controllers.ChangeCapabilityHandler)
		//删除能力面板
		v1.POST("/capability/delete", controllers.DeleteCapabilityHandler)
	}
}
