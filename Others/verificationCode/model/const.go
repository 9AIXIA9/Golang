package model

const (
	ErrForPhoneFormate         = "手机号码格式错误"
	ErrForVerTimesHasExceed    = "今天获取验证码次数已经超过五次"
	ErrForVerIntervalNotEnough = "五分钟内不可以再次获取验证码"
	ErrorForCodeHasExpired     = "验证码已过期"
)

const (
	OpenFileError = "打开文件失败"
)

type PhoneDetail struct {
	ValidDuration int64  //时间戳 用来判断是否过期
	Interval      int64  //时间戳 用来判断是否在五分钟内
	Times         int8   //次数
	Code          string //验证码
}
