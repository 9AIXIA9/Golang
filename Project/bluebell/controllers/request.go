package controllers

import (
	"bluebell/myerrors"
	"errors"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey = "userID"
)

// GetCurrentUser 获取当前登录用户id
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = myerrors.UserNotLogin
	}
	userID, ok = uid.(int64)
	if !ok {
		err = myerrors.UserNotLogin
	}
	return
}

// GetPageInfo 获取页面信息
func GetPageInfo(c *gin.Context) (pageNum, pageSize int64, err error) {
	//获取分页参数
	pageNumStr := c.Query("page_num")
	pageSizeStr := c.Query("page_size")

	pageNum, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		zap.L().Error(errors.Join(myerrors.PostPageGet, err).Error())
		return
	}
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		zap.L().Error(errors.Join(myerrors.PostPageGet, err).Error())
	}
	return
}
