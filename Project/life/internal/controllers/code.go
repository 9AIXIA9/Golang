package controllers

import "net/http"

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeServerBusy
	CodeUserExist
	CodeUserNotExist
	CodeWrongPassword
	CodeNeedLogin
	CodeInvalidToken
	CodeInvalidParam
	CodeRateLimit
	CodeTooFrequent
	CodeCapExist
	CodeCapNotExist
	CodeNoneCap
	CodeUserCapNotExist
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:         "成功",
	CodeServerBusy:      "服务繁忙",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeCapExist:        "用户该能力值已存在",
	CodeCapNotExist:     "用户该能力值不存在",
	CodeNoneCap:         "用户不存在能力值",
	CodeUserCapNotExist: "用户能力为空",
	CodeWrongPassword:   "账号或密码错误",
	CodeNeedLogin:       "用户未登录",
	CodeInvalidToken:    "令牌错误",
	CodeInvalidParam:    "参数错误",
	CodeRateLimit:       "访问限速，请稍后再试",
	CodeTooFrequent:     "频率过高，请稍后再试",
	http.StatusNotFound: "页面未定义",
}

func (c ResCode) Msg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
