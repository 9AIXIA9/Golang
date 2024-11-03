package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeInvalidUserOrPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeCommunityNotExist
	CodeNonePage
	CodePostVoteOverdue
	CodePostVoteNotChange
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:               "Success",
	CodeInvalidParam:          "输入有误字符",
	CodeUserExist:             "用户已存在",
	CodeInvalidUserOrPassword: "输入的用户或密码错误",
	CodeServerBusy:            "服务繁忙",
	CodeNeedLogin:             "需要登录",
	CodeInvalidToken:          "无效Token",
	CodeCommunityNotExist:     "访问社区不存在",
	CodeNonePage:              "访问页面不存在",
	CodePostVoteOverdue:       "帖子投票时间过期",
	CodePostVoteNotChange:     "您已投过此类型的票",
}

func (c ResCode) Msg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
