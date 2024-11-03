package logic

import (
	"learn/dao/mysql"
	"learn/model"
)

// Login 用于检测是否存在该数据
func Login(u model.User) bool {
	return mysql.Login(u)
}

// CheckExist 用于检测用户是否存在
func CheckExist(u model.User) bool {
	return mysql.CheckExist(u)
}

// UpLoad 用于上传用户数据到数据库
func UpLoad(u model.User) {
	mysql.UpLoad(u)
}
