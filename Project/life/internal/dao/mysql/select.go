package mysql

import (
	"database/sql"
	"errors"
	"life/pkg/models"

	"go.uber.org/zap"
)

// QueryUserIDPasswordByUserName 通过用户名字查询用户ID和密码
func QueryUserIDPasswordByUserName(username string) (userID int64, password string, err error) {
	sqlStr := `SELECT password,user.user_id from user where username=?`
	err = db.QueryRow(sqlStr, username).Scan(&password, &userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Error("mysql: query userID and password failed,user not exist",
				zap.String("username", username),
				zap.Error(err))
			return 0, "", ErrorDataNil
		}
		zap.L().Error("mysql query userID and password failed",
			zap.String("username", username),
			zap.Error(err))
		return 0, "", errorOperate
	}
	//zap.L().Info("db", zap.String("user name", username), zap.String("password", password), zap.Int64("userID", userID))
	return userID, password, nil
}

// QueryCapabilityByUserID 通过用户ID查询用户能力
func QueryCapabilityByUserID(userID int64) ([]*models.CapInfo, error) {
	const sqlStr = `SELECT capability_name, capability_score
		FROM capability
		WHERE user_id = ?
		ORDER BY update_time`
	var capabilities []*models.CapInfo
	err := db.Select(&capabilities, sqlStr, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Error("user dont have capability",
				zap.Int64("userID", userID),
				zap.Error(err))
			return nil, ErrorDataNil
		}
		zap.L().Error("mysql select user capability failed",
			zap.Int64("userID", userID),
			zap.Error(err))
		return nil, err
	}
	return capabilities, nil
}

// QueryCapabilityByIDName 通过用户ID和能力名称查询能力数值
func QueryCapabilityByIDName(userID int64, name string) (score int16, err error) {
	const sqlStr = `
        SELECT capability_score
        FROM capability
        WHERE user_id = ? AND capability_name = ?
        ORDER BY update_time DESC
        LIMIT 1`
	if err = db.QueryRow(sqlStr, userID, name).Scan(&score); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Error("mysql: query capability score failed,user capability not exist",
				zap.Int64("userID", userID),
				zap.String("capabilityName", name),
				zap.Error(err))
			return 0, ErrorDataNil
		}
		zap.L().Error("mysql query user capability failed",
			zap.Int64("userID", userID),
			zap.String("capabilityName", name),
			zap.Error(err))
		return 0, errorOperate
	}
	return score, nil
}
