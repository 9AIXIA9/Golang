package controllers

import (
	"bluebell/logic"
	"bluebell/myerrors"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetCommunityHandler 处理社区访问请求的函数
func GetCommunityHandler(c *gin.Context) {
	//查询到所有社区(community_id,community_name)以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error(err.Error())
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	return
}

// GetCommunityDetailHandler 处理特定id社区访问请求的函数
func GetCommunityDetailHandler(c *gin.Context) {
	//1.获取社区id
	idStr := c.Param("community_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.用id查询社区数据
	data, err := logic.GetCommunityDetailByID(id)
	if err != nil {
		zap.L().Error(err.Error())
		if errors.Is(err, myerrors.CommunityInvalidID) {
			ResponseError(c, CodeCommunityNotExist)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	//3.返回响应
	ResponseSuccess(c, data)
	return
}
