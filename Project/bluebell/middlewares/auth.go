package middlewares

import (
	"bluebell/controllers"
	"bluebell/dao/redis"
	"bluebell/myerrors"
	"bluebell/pkg/jwt"
	"bluebell/settings"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//客户端携带Token有三种方式
		//1.放在请求头 2.放在请求头 3.放在URI
		//这里假设Token放在Header的Authorization中，并使用Bearer开头
		//Authorization:Bearer xxx.xxx.xx
		//这里的具体实现方式要依据你的实际业务情况而定
		authHeader := c.Request.Header.Get(settings.Config.TokenLocation)
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			zap.L().Error(myerrors.UserNotLogin.Error())
			c.Abort()
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == settings.Config.TokenHeader) {
			zap.L().Error(myerrors.UserNotLogin.Error())
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			c.Abort()
			return
		}
		//part[1]是获取到的tokenString，我们使用之前定义好的JWT的函数来解析它
		token := parts[1]
		mc, err := jwt.ParseToken(token)
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			zap.L().Error(myerrors.InvalidToken.Error())
			c.Abort()
			return
		}

		tokenDB, err := redis.GetTokenByUserID(mc.UserID)
		//处理错误
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			zap.L().Error(myerrors.InvalidToken.Error())
			c.Abort()
			return
		} else if !(tokenDB == token) {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			zap.L().Error(myerrors.InvalidToken.Error())
			c.Abort()
			return
		}
		//将当前请求的userID信息保存到请求的上下文c上
		c.Set(controllers.CtxUserIDKey, mc.UserID)
		c.Next() //后续的处理函数可以用c.Get("userID")来获取当前请求的用户信息
	}
}
