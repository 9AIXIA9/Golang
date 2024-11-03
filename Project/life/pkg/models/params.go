package models

import "strings"

//uto

// ParamSignUp 用户注册传输参数规范
type ParamSignUp struct {
	Username   string `json:"username" binding:"required,min=2,max=10,alphanum"`
	Password   string `json:"password" binding:"required,min=8,max=20"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 用户登录传输参数规范
type ParamLogin struct {
	Username string `json:"username" binding:"required,min=2,max=10,alphanum"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

// ParamUpdateInfo 用户更新信息参数规范
type ParamUpdateInfo struct {
	Email  string `json:"email" binding:"required,email"`
	Gender int8   `json:"gender" binding:"oneof= 1 0 -1"` //1为男 -1为女 0为未知
}

// ParamSetCapability 用户建立能力面板
type ParamSetCapability struct {
	Name       string `json:"capability_name" binding:"required"`              //能力名字
	BasisScore int16  `json:"capability_basic_score" binding:"gte=0,lte=1000"` //能力基础数值
}

// ParamUpdateCapability 用户更新能力面板
type ParamUpdateCapability struct {
	Name   string `json:"capability_name" binding:"required"`               //能力名字
	Change int16  `json:"capability_score_change" binding:"gte=0,lte=1000"` //能力数值
}

// ParamDeleteCapability 用户删除能力面板
type ParamDeleteCapability struct {
	Name string `json:"capability_name" binding:"required"` //能力名字
}

// RemoveTopStruct 去除提示信息的结构体名称
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		//strings.Index(field, ".") 查找字符串中第一个点号的位置
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
