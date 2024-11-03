package controllers

//定义controllers层级公用函数错误
import "errors"

var (
	errorGetUserID = errors.New("get userID in context failed")
)
