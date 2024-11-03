package api

import (
	"life/internal/settings"
	"life/pkg/utils/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var r *gin.Engine

// SetupRoutes 建立路由
func SetupRoutes() *gin.Engine {
	Init()
	//建立 v1系 路由
	setUpV1()
	//未定义页面路由
	setUpNoRoute()
	return r
}

func Init() {
	// 如果配置的模式无效，设置默认值
	mode := settings.Config.App.Mode
	if mode != gin.DebugMode && mode != gin.ReleaseMode && mode != gin.TestMode {
		zap.L().Error("gin mode set failed,falling back to DebugMode")
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery())
}
