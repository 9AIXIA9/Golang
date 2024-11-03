package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/myerrors"
	"errors"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// PostVoteHandler 用于处理帖子投票路由
func PostVoteHandler(c *gin.Context) {
	// 获取参数——用户投的什么票帖子id
	p := new(models.ParamVoteData)
	if err := c.ShouldBind(p); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			errData := models.RemoveTopStruct(errs.Translate(translator))
			zap.L().Error(myerrors.InvalidParam.Error(), zap.Error(err), zap.Any("details", errData))
			ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		} else {
			zap.L().Error(myerrors.InvalidParam.Error(), zap.Error(err))
			ResponseError(c, CodeInvalidParam)
		}
		return
	}

	// 获取参数用户id
	userID, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error(myerrors.InvalidToken.Error(), zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 逻辑处理
	if err = logic.PostVote(p, userID); err != nil {
		zap.L().Error(myerrors.PostVote.Error(), zap.Error(err))
		switch {
		case errors.Is(err, myerrors.PostVoteOverdue):
			ResponseError(c, CodePostVoteOverdue)
		case errors.Is(err, myerrors.PostVoteNotChange):
			ResponseError(c, CodePostVoteNotChange)
		default:
			ResponseError(c, CodeServerBusy)
		}
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
}
