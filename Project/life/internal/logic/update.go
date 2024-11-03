package logic

import (
	"errors"
	"life/internal/dao/mysql"
	"life/internal/dao/redis"

	"go.uber.org/zap"
)

// UpdateCapability 更新用户能力值
func UpdateCapability(userID int64, name string, score int16) error {
	//更新mysql数据
	if err := mysql.UpdateCapability(userID, name, score); err != nil {
		return err
	}
	return nil
}

// DeleteCapCache 删除数据库缓存(redis)
func DeleteCapCache(userID int64) error {
	//删除redis数据
	if err := redis.DeleteCapByID(userID); err != nil {
		if errors.Is(err, redis.ErrorNoneCap) {
			zap.L().Error("redis not have user capability", zap.Int64("userID", userID))
			return ErrorUserNoneCap
		}
		return err
	}
	return nil
}
