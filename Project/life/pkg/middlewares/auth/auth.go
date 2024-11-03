package auth

//
import (
	"errors"
	"life/internal/controllers"
	"life/internal/dao/redis"
	"life/internal/settings"
	"life/pkg/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	errorNoToken    = errors.New("user not pass token")
	errorFalseToken = errors.New("user pass false token")
)

// JWTMiddleware 基于JWT的认证中间件，完成
func JWTMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(settings.Config.TokenLocation)
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			zap.L().Error("auth header null", zap.Error(errorNoToken))
			c.Abort()
			return
		}

		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == settings.Config.TokenHeader) {
			zap.L().Error("invalid auth header", zap.Error(errorFalseToken))
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			c.Abort()
			return
		}

		//part[1]是获取到的tokenString，我们使用之前定义好的JWT的函数来解析它
		token := parts[1]
		mc, err := jwt.ParseToken(token)
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			zap.L().Error("parse token failed", zap.Error(err))
			c.Abort()
			return
		}

		tokenDB, err := redis.GetTokenByUserID(mc.UserID)
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			zap.L().Error("get redis token failed", zap.Error(err))
			c.Abort()
			return
		} else if !(tokenDB == token) {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			zap.L().Error("false token,redis token != passed token")
			c.Abort()
			return
		}
		//将当前请求的userID信息保存到请求的上下文c上
		c.Set(controllers.CtxUserIDKey, mc.UserID)
		c.Next() //后续的处理函数可以用c.Get("userID")来获取当前请求的用户信息
	}
}
