package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn/logic"
	"learn/model"
	"net/http"
	"strconv"
)

// 用户登录状态
var userLoadStatus = false

// 用户注册状态
var userEnrollStatus = false

// LoginUser 用于用户登录
func LoginUser(ctx *gin.Context) {
	//ctx.HTML(http.StatusOK, "login.html", nil)
	//获取这里的基本讯息，id,password
	var u model.User
	//var err error
	u.Id = 111
	//phone := ctx.PostForm("phone")
	//u.Password = ctx.PostForm("password")
	//u.Phone, err = strconv.Atoi(phone)
	//if err = ctx.ShouldBind(&u); err != nil {
	//	fmt.Println(err)
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"msg": err,
	//	})
	//	return
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": u,
	})
	//fmt.Printf("%#v\n", u)
	//
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//	return
	//}
	//ok := logic.Login(u)
	//if !ok {
	//	fmt.Printf("账户或密码错误!\n")
	//	return
	//}
	//userLoadStatus = true
}

// EnrollUser 用于用户注册
func EnrollUser(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "enroll.html", nil)
	//获取这里的基本讯息，id,password,name,phone
	var u model.User
	var err error
	phoneString := ctx.PostForm("phone")
	u.Password = ctx.PostForm("password")
	u.Name = ctx.PostForm("name")
	u.Phone, err = strconv.Atoi(phoneString)
	fmt.Printf("%#v\n", u)
	//phoneString 拿不到
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	ok := logic.CheckExist(u)
	if !ok {
		fmt.Printf("用户已存在!\n")
		return
	}
	logic.UpLoad(u)
	userEnrollStatus = true
}

// LoginUserPage 用户登录后跳转页面
func LoginUserPage(ctx *gin.Context) {
	if userLoadStatus {
		ctx.HTML(200, "index.html", nil)
		return
	}
	ctx.HTML(http.StatusForbidden, "login.html", nil)
}

// EnrollUserPage 注册后跳转页面
func EnrollUserPage(ctx *gin.Context) {
	if userEnrollStatus {
		ctx.HTML(http.StatusOK, "login.html", nil)
		return
	}
	ctx.HTML(http.StatusForbidden, "enroll.html", nil)
}
