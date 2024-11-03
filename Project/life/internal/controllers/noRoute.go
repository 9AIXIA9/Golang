package controllers

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// NoRouteHandler 处理未定义页面
func NoRouteHandler(c *gin.Context) {
	//记录信息
	zap.L().Info("user visit undefined page")
	//返回响应
	ResponseError(c, http.StatusNotFound)
}
