package controllers

import (
	"errors"
	"life/internal/logic"
	"life/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// SignUpHandler 用户注册页面处理
func SignUpHandler(c *gin.Context) {
	//获取用户传入参数
	var p models.ParamSignUp
	if err := c.ShouldBind(&p); err != nil {
		//判断err是不是validator.validationErrors类型
		var typeErr validator.ValidationErrors
		ok := errors.As(err, &typeErr)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, models.RemoveTopStruct(typeErr.Translate(translator)))
		return
	}
	//逻辑处理
	userID, token, err := logic.Signup(p.Username, p.Password)
	if err != nil {
		if errors.Is(err, logic.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	userInfo := models.UserInfo{
		Username: p.Username,
		UserID:   userID,
		Token:    token,
	}
	ResponseSuccess(c, userInfo)
}

// LoginHandler 用户登录页面处理
func LoginHandler(c *gin.Context) {
	//获取用户传入参数
	var p models.ParamLogin
	if err := c.ShouldBind(&p); err != nil {
		//判断err是不是validator.validationErrors类型
		var typeErr validator.ValidationErrors
		ok := errors.As(err, &typeErr)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, models.RemoveTopStruct(typeErr.Translate(translator)))
		return
	}

	//逻辑处理
	userID, token, err := logic.Login(p.Username, p.Password)
	if err != nil {
		if errors.Is(err, logic.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, logic.ErrorWrongPassword) {
			ResponseError(c, CodeWrongPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	userInfo := models.UserInfo{
		Username: p.Username,
		UserID:   userID,
		Token:    token,
	}
	ResponseSuccess(c, userInfo)
}

// UpdateInfoHandler 用于处理更新个人信息
func UpdateInfoHandler(c *gin.Context) {
	//获取参数
	var p models.ParamUpdateInfo
	if err := c.ShouldBind(&p); err != nil {
		//判断err是不是validator.validationErrors类型
		var typeErr validator.ValidationErrors
		ok := errors.As(err, &typeErr)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, models.RemoveTopStruct(typeErr.Translate(translator)))
		return
	}
	//获取当前用户信息
	userID, err := GetUserID(c)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	//更改当前用户信息
	if err := logic.UpdateInfo(userID, p.Email, p.Gender); err != nil {
		if errors.Is(err, logic.ErrorTooFrequent) {
			ResponseError(c, CodeTooFrequent)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
