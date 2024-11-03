package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	CtxUserIDKey = "userID"
)

// GetUserID 从上下文中取出用户ID
func GetUserID(c *gin.Context) (userID int64, err error) {
	//从 Context取 userID
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		zap.L().Debug("get user ID in context failed")
		return 0, errorGetUserID
	}
	//转化uid成int64
	userID, ok = uid.(int64)
	if !ok {
		zap.L().Error("convert uid to int64 failed")
		return 0, errorGetUserID
	}
	return
}
