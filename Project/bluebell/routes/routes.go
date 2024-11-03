package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("api/v1")
	//注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	//登录业务路由
	v1.POST("/login", controllers.LoginHandler)
	//应用JWT认证中间件
	v1.Use(middlewares.JWTAuthMiddleware(), middlewares.RateLimitMiddleware(2*time.Second, 1))
	//登录后才能访问的页面
	{
		v1.GET("/community", controllers.GetCommunityHandler)
		// /:指定路径参数
		v1.GET("/community/:community_id", controllers.GetCommunityDetailHandler)
		v1.GET("/community/post", controllers.GetCommunityPostListHandler)

		v1.POST("/post", controllers.PostCreateHandler)
		v1.GET("/posts", controllers.GetPostListHandler)
		v1.GET("/posts2", controllers.GetPostListHandler2)

		v1.POST("/post/:id", controllers.PostGetHandler)

		v1.POST("/vote", controllers.PostVoteHandler)

	}
	//未定义页面路由
	r.NoRoute(controllers.NoRouteHandler)
	return r
}
