package models

import (
	"strings"
)

const (
	OrderTime  = "time"
	OrderScore = "score"
)

//定义请求的参数结构体

// ParamSignUp 注册参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录参数结构体
type ParamLogin struct {
	UserID   int64  `json:"user_id,string"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 帖子投票参数结构体
type ParamVoteData struct {
	//UserID从token中获取
	PostID    int64 `json:"post_id,string" binding:"required"` //帖子id
	Direction int8  `json:"direction,string" binding:"oneof= 1 0 -1"`
	//去掉required的原因是 他会导致无法填零值
	//投赞成(1)还是反对(-1)取消(0)
}

// ParamPostList 获取帖子查询参数
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"` //帖子id
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
	//去掉required的原因是 他会导致无法填零值
	//投赞成(1)还是反对(-1)取消(0)
}

// ParamCommunityPostList 获取帖子查询参数
type ParamCommunityPostList struct {
	*ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}

// RemoveTopStruct 去除提示信息的结构体名称
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
