package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code ResCode     `json:"code,omitempty"`
	Msg  interface{} `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// ResponseError 返回错误信息
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusInternalServerError, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 返回错误及信息
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusInternalServerError, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 返回成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

// ResponseRateLimit 返回速率限制
func ResponseRateLimit(c *gin.Context) {
	c.JSON(http.StatusTooManyRequests, &ResponseData{
		Code: CodeRateLimit,
		Msg:  CodeRateLimit.Msg(),
		Data: nil,
	})
}
