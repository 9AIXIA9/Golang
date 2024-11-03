package controllers

import (
	"bluebell/myerrors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NoRouteHandler NoRoute 处理无法访问
func NoRouteHandler(c *gin.Context) {
	//返回响应
	zap.L().Error(myerrors.NoRoute.Error())
	ResponseError(c, CodeNonePage)
}
