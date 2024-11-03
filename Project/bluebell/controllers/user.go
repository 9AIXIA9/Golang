package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/myerrors"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1.参数获取和校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error(myerrors.InvalidParam.Error(), zap.Error(err))
		//判断err是不是validator.validationErrors类型
		var vErr validator.ValidationErrors
		ok := errors.As(err, &vErr)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, models.RemoveTopStruct(vErr.Translate(translator)))
		return
	}
	//2.业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error(myerrors.SignUp.Error(), zap.Error(err))
		if errors.Is(err, myerrors.UserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求
func LoginHandler(c *gin.Context) {
	//1.参数获取和校验
	var p models.ParamLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error(myerrors.InvalidParam.Error(), zap.Error(err))
		//判断err是不是validator.validationErrors类型
		var vErr validator.ValidationErrors
		ok := errors.As(err, &vErr)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, models.RemoveTopStruct(vErr.Translate(translator)))
		return
	}
	//2.业务处理
	user, err := logic.Login(&p)
	if errors.Is(err, myerrors.UserExist) {
		ResponseError(c, CodeInvalidUserOrPassword)
		zap.L().Error(err.Error())
		return
	} else if errors.Is(err, myerrors.UserWrongPassword) {
		ResponseError(c, CodeInvalidUserOrPassword)
		zap.L().Error(err.Error())
		return
	} else if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error(myerrors.Login.Error(), zap.Error(err))
		return
	}
	//3.返回响应
	ResponseSuccess(c, user)
}
