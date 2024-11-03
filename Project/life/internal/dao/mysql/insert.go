package mysql

import "go.uber.org/zap"

// InsertUser 在数据库中插入用户
func InsertUser(userID int64, username, password string) (err error) {
	//执行SQL语句入库
	sqlStr := `INSERT INTO user(user_id, username, password) values	(?,?,?)`
	_, err = db.Exec(sqlStr, userID, username, password)
	if err != nil {
		zap.L().Error("mysql insert user failed", zap.Error(err))
		zap.L().Error("insert capability failed",
			zap.Int64("user_id", userID),
			zap.String("username", username),
			zap.Error(err))
		return err
	}
	return
}

// InsertCapability 传入能力值
func InsertCapability(userID int64, capName string, basicScore int16) (err error) {
	sqlStr := `INSERT INTO capability(user_id, capability_name, capability_score) VALUES (?,?,?)`
	_, err = db.Exec(sqlStr, userID, capName, basicScore)
	if err != nil {
		zap.L().Error("insert capability failed",
			zap.Int64("user_id", userID),
			zap.String("capability_name", capName),
			zap.Int("score", int(basicScore)),
			zap.Error(err))
		return err
	}
	return nil
}
