package mysql

import (
	"bluebell/models"
	"bluebell/myerrors"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

//把每一步数据库操作封装成函数
//待logic层根据业务需求调用

const secret = "XIA"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `SELECT COUNT(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return
	}
	if count > 0 {
		err = myerrors.UserExist
	}
	return
}

// InsertUser 在数据库中插入用户
func InsertUser(user *models.User) (err error) {
	//对用户原始密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `INSERT INTO user(user_id, username, password) values	(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// QueryUserPasswordID 通过名字查询用户密码和ID
func QueryUserPasswordID(username string) (password string, userID int64, err error) {
	sqlStr := `SELECT password,user.user_id from user where username=?`
	if err = db.QueryRow(sqlStr, username).Scan(&password, &userID); errors.Is(err, sql.ErrNoRows) {
		err = myerrors.UserNotExist
	}
	return
}

// CheckPassword 核对输入密码是否正确
func CheckPassword(originPassword string, password string) (err error) {
	if encryptPassword(originPassword) != password {
		err = myerrors.UserWrongPassword
	}
	return
}

// encryptPassword 对用户密码进行加密
func encryptPassword(originPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(originPassword)))
}

// GetUserByID 通过ID查询用户
func GetUserByID(userID int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `SELECT username from user where user_id=?`
	err = db.Get(user, sqlStr, userID)
	if errors.Is(err, sql.ErrNoRows) {
		err = myerrors.UserNotExist
	}
	return
}
