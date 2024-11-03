package mysql

import "go.uber.org/zap"

// UpdateUserInfoByID 通过用户ID更新用户信息
func UpdateUserInfoByID(userID int64, email string, gender int8) (err error) {
	// 执行SQL语句更新
	sqlStr := `UPDATE user SET email=?, gender=? WHERE user_id=?`
	result, err := db.Exec(sqlStr, email, gender, userID)
	if err != nil {
		zap.L().Error("mysql update user information failed",
			zap.Int64("user_id", userID),
			zap.String("email", email),
			zap.Int("gender", int(gender)),
			zap.Error(err))
		return err
	}
	// 检查是否有记录被更新
	rows, err := result.RowsAffected()
	if err != nil {
		zap.L().Error("mysql get RowsAffected failed", zap.Error(err))
		return err
	}
	if rows == 0 {
		zap.L().Error("mysql user not found or data not change")
		return errorOperate
	}
	return nil
}

// UpdateCapability 更新个人能力
func UpdateCapability(userID int64, name string, change int16) (err error) {
	// 执行SQL语句更新
	sqlStr := `UPDATE capability SET capability_score=? WHERE user_id=? AND capability_name=?`
	result, err := db.Exec(sqlStr, change, userID, name)
	if err != nil {
		zap.L().Error("mysql update user information failed",
			zap.Int64("user_id", userID),
			zap.String("capability_name", name),
			zap.Int16("capability_score_change", change),
			zap.Error(err))
		return err
	}
	// 检查是否有记录被更新
	rows, err := result.RowsAffected()
	if err != nil {
		zap.L().Error("mysql get RowsAffected failed",
			zap.Int64("user_id", userID),
			zap.String("capability_name", name),
			zap.Int16("capability_score_change", change),
			zap.Error(err))
		return err
	}
	if rows == 0 {
		zap.L().Error("mysql user capability not found or data not change",
			zap.Int64("user_id", userID),
			zap.String("capability_name", name),
			zap.Int16("capability_score_change", change))
		return errorOperate
	}
	return nil
}
