package mysql

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"go.uber.org/zap"
)

// CheckUserExistByName 通过名字检查用户是否存在
func CheckUserExistByName(userName string) (exist bool, err error) {
	sqlStr := `SELECT COUNT(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, userName); err != nil {
		zap.L().Error("mysql get user by name failed",
			zap.String("username", userName),
			zap.Error(err))
		return false, errorOperate
	}
	if count > 0 {
		exist = true
	}
	return exist, nil
}

// CheckPassword 核对密码是否正确
func CheckPassword(loginPassword, dbPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(loginPassword)) //验证（对比）
}

// CheckUserUpdateTime 检查更新时间是否符合标准
func CheckUserUpdateTime(userID int64) (bool, error) {
	var (
		lastUpdate time.Time
		createTime time.Time
	)
	const sqlStr = `SELECT update_time,create_time FROM user WHERE user_id = ?`
	if err := db.QueryRow(sqlStr, userID).Scan(&lastUpdate, &createTime); err != nil {
		zap.L().Error("mysql query user update time failed",
			zap.Int64("user_id", userID),
			zap.Error(err))
		return false, errorOperate
	}
	//第一次修改
	if createTime == lastUpdate {
		//第一次更改信息
		return true, nil
	}
	//判断是否超时
	if time.Since(lastUpdate) < userInterval {
		return true, nil
	}

	return false, nil
}

// CheckUserCapExist 检查用户能力是否存在
func CheckUserCapExist(userID int64, capName string) (exist bool, err error) {
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM capability WHERE user_id=? AND capability_name=?)", userID, capName).Scan(&exist)
	if err != nil {
		zap.L().Error("check capability existence failed",
			zap.Int64("userID", userID),
			zap.String("capabilityName", capName),
			zap.Error(err))
		return true, errorOperate
	}
	return exist, nil
}

// CheckCapUpdateTime 检查能力值的更新时间是否符合标准
func CheckCapUpdateTime(userID int64, capName string) (bool, error) {
	var (
		lastUpdate time.Time
		createTime time.Time
	)

	const sqlStr = `SELECT update_time,create_time FROM capability WHERE user_id=? AND capability_name= ?`

	if err := db.QueryRow(sqlStr, userID, capName).Scan(&lastUpdate, &createTime); err != nil {
		zap.L().Error("mysql query user update time failed",
			zap.Int64("user_id", userID),
			zap.String("capability_name", capName),
			zap.Error(err))
		return false, errorOperate
	}

	if createTime == lastUpdate {
		//第一次更改信息
		return true, nil
	}
	//判断是否超时
	if time.Since(lastUpdate) < capInterval {
		return true, nil
	}

	return false, nil
}
