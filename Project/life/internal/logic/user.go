package logic

import (
	"errors"
	"life/internal/dao/mysql"
	"life/internal/dao/redis"
	"life/pkg/utils/jwt"
	"life/pkg/utils/snowflake"

	"go.uber.org/zap"
)

// Signup 处理注册逻辑
func Signup(username, password string) (userID int64, token string, err error) {
	//查询用户名称是否存在
	var exist bool
	exist, err = mysql.CheckUserExistByName(username)
	if err != nil { //mysql出错
		return
	}
	//用户已存在
	if exist {
		zap.L().Error(ErrorUserExist.Error(),
			zap.String("username", username))
		return 0, "", ErrorUserExist
	}
	//通过雪花算法生成用户ID
	userID = snowflake.GenID()
	//通过encrypt实现密码加密
	ePassword, err := encrypt(password)
	if err != nil {
		return 0, "", err
	}
	//存储数据库
	if err = mysql.InsertUser(userID, username, ePassword); err != nil {
		zap.L().Error("update user information database failed")
		return 0, "", err
	}
	//生成令牌
	token, err = jwt.GenToken(userID, username)
	if err != nil {
		zap.L().Error("generate token failed",
			zap.String("username", username),
			zap.Int64("userID", userID),
			zap.Error(err))
		return 0, "", err
	}
	//将token存储在redis中
	err = redis.StoreToken(userID, token)
	if err != nil {
		return 0, "", err
	}
	return userID, token, nil
}

// Login 处理登录逻辑
func Login(username, password string) (userID int64, token string, err error) {
	//查询用户是否存在
	var (
		exist      bool
		dbPassword string
	)
	exist, err = mysql.CheckUserExistByName(username)
	if err != nil { //mysql出错
		return 0, "", err
	}
	if !exist {
		zap.L().Error(ErrorUserNotExist.Error(), zap.String("username", username))
		return 0, "", ErrorUserNotExist
	}
	//查询密码和ID
	userID, dbPassword, err = mysql.QueryUserIDPasswordByUserName(username)
	if err != nil {
		if errors.Is(err, mysql.ErrorDataNil) {
			return 0, "", ErrorUserNotExist
		}
		return 0, "", err
	}
	//核对密码是否正确
	err = mysql.CheckPassword(password, dbPassword)
	//zap.L().Info("compare password", zap.String("password", password), zap.String("dbPassword", dbPassword))
	if err != nil {
		zap.L().Error(ErrorWrongPassword.Error(), zap.String("username", username))
		return 0, "", ErrorWrongPassword
	}
	//生成令牌
	token, err = jwt.GenToken(userID, username)
	if err != nil {
		zap.L().Error("generate token failed",
			zap.String("username", username),
			zap.Int64("userID", userID),
			zap.Error(err))
		return 0, "", err
	}
	//将token存储在redis中
	err = redis.StoreToken(userID, token)
	if err != nil {
		return 0, "", err
	}
	return
}

// UpdateInfo 处理更新信息逻辑
func UpdateInfo(userID int64, email string, gender int8) (err error) {
	//验证更新时间是否合理
	var ok bool
	ok, err = mysql.CheckUserUpdateTime(userID)
	if err != nil { //mysql出错
		return err
	}
	if !ok {
		zap.L().Error(ErrorTooFrequent.Error(),
			zap.Int64("userID", userID))
		return ErrorTooFrequent
	}
	//更新信息
	err = mysql.UpdateUserInfoByID(userID, email, gender)
	if err != nil { //mysql出错
		return err
	}
	return nil
}
