package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// 业务逻辑的代码

// SignUp //用于注册用户
func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//2.生成UID
	userID := snowflake.GenID()
	//构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.保存进数据库
	return mysql.InsertUser(user)
}

// Login 用于登录用户
func Login(p *models.ParamLogin) (user *models.User, err error) {
	//1.获取用户密码和UserID
	var encryptPassword string
	if encryptPassword, p.UserID, err = mysql.QueryUserPasswordID(p.Username); err != nil {
		return nil, err
	}
	//2.核对密码
	if err = mysql.CheckPassword(p.Password, encryptPassword); err != nil {
		return nil, err
	}

	//3.成功登录
	//生成token令牌
	token, err := jwt.GenToken(p.UserID, p.Username)
	if err != nil {
		return nil, err
	}
	user = &models.User{
		UserID:   p.UserID,
		Username: p.Username,
		Token:    token,
	}
	return
}
