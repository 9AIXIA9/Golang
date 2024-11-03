package sync

import (
	"life/internal/dao/mysql"
	"life/internal/dao/redis"
	"strconv"

	"go.uber.org/zap"
)

// UpdateCapability 更新能力值
func UpdateCapability() error {
	zap.L().Info("start update capability")
	//获取redis中所有值
	allUserCaps, err := redis.QueryAllUser()
	if err != nil {
		return err
	}
	zap.L().Info("update", zap.Any("all user caps", allUserCaps))
	for _, userCap := range allUserCaps {
		//更新能力值
		for name, score := range userCap.Capabilities {
			s, err := strconv.ParseInt(score, 10, 16)
			if err != nil {
				zap.L().Error("user update capability failed",
					zap.Int64("userID", userCap.UserID),
					zap.String("capName", name),
					zap.Error(err))
				continue
			}
			//取出mysql能力值
			dScore, err := mysql.QueryCapabilityByIDName(userCap.UserID, name)
			if err != nil {
				zap.L().Error("user update capability failed",
					zap.Int64("userID", userCap.UserID),
					zap.String("capName", name),
					zap.Error(err))
				continue
			}
			//更新mysql能力值
			err = mysql.UpdateCapability(userCap.UserID, name, dScore+int16(s))
			if err != nil {
				zap.L().Error("user update capability failed",
					zap.Int64("userID", userCap.UserID),
					zap.String("capName", name),
					zap.Error(err))
				continue
			}
		}
		//删除缓存
		err := redis.DeleteCapByID(userCap.UserID)
		if err != nil {
			zap.L().Error("user delete capability failed",
				zap.Int64("userID", userCap.UserID),
				zap.Error(err))
			continue
		}
	}
	zap.L().Info("finish update capability")
	return nil
}
