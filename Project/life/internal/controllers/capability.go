package controllers

import (
	"errors"
	"life/internal/logic"
	"life/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// InsertCapabilityHandler 建立个人能力
func InsertCapabilityHandler(c *gin.Context) {
	//获取参数
	//获取用户传入参数
	var p models.ParamSetCapability
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
	//逻辑处理
	if err := logic.SetUpCapability(userID, p.Name, p.BasisScore); err != nil {
		if errors.Is(err, logic.ErrorCapExist) {
			ResponseError(c, CodeCapExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

// GetCapabilityHandler 获取个人能力表格
func GetCapabilityHandler(c *gin.Context) {
	//获取当前用户信息
	userID, err := GetUserID(c)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	//逻辑处理
	data, err := logic.GetCapabilityByID(userID)
	if err != nil {
		if errors.Is(err, logic.ErrorUserNoneCap) {
			ResponseError(c, CodeNoneCap)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// ChangeCapabilityHandler 更新个人能力
func ChangeCapabilityHandler(c *gin.Context) {
	//获取参数
	//获取用户传入参数
	var p models.ParamUpdateCapability
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
	//逻辑处理
	if err := logic.ChangeCapability(userID, p.Name, p.Change); err != nil {
		if errors.Is(err, logic.ErrorCapNotExist) {
			ResponseError(c, CodeCapNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

func DeleteCapabilityHandler(c *gin.Context) {
	//获取参数
	//获取用户传入参数
	var p models.ParamDeleteCapability
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
	//逻辑处理
	if err := logic.DeleteCapByName(userID, p.Name); err != nil {
		if errors.Is(err, logic.ErrorCapNotExist) {
			ResponseError(c, CodeCapNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
