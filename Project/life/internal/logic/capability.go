package logic

import (
	"errors"
	"life/internal/dao/mysql"
	"life/internal/dao/redis"
	"life/pkg/models"

	"go.uber.org/zap"
)

// SetUpCapability 处理建立能力逻辑
func SetUpCapability(userID int64, name string, basicScore int16) (err error) {
	//判断用户能力值是否存在
	var exist bool
	exist, err = mysql.CheckUserCapExist(userID, name)
	if err != nil {
		return err
	}
	if exist {
		zap.L().Error("user capability have existed",
			zap.Int64("user_id", userID),
			zap.String("capability_name", name),
			zap.Int16("capability_basic_score", basicScore),
		)
		return ErrorCapExist
	}
	//插入能力值
	err = mysql.InsertCapability(userID, name, basicScore)
	if err != nil {
		return err
	}
	err = redis.StoreCapScore(userID, name, basicScore)
	if err != nil {
		return err
	}
	return nil
}

// GetCapabilityByID 处理通过userID查询能力值逻辑
func GetCapabilityByID(userID int64) ([]*models.CapInfo, error) {
	changeMap, err := redis.GetAllCapScore(userID)
	if err != nil && !errors.Is(err, redis.ErrorNoneCap) {
		return nil, err
	}
	capabilities, err := mysql.QueryCapabilityByUserID(userID)
	if err != nil {
		if errors.Is(err, mysql.ErrorDataNil) {
			zap.L().Error(ErrorUserNoneCap.Error(),
				zap.Int64("userID", userID))
			return nil, ErrorUserNoneCap
		}
		return nil, err
	}
	//判断用户能力面板是否为空
	if len(capabilities) == 0 {
		zap.L().Error(ErrorUserNoneCap.Error(), zap.Int64("user_id", userID))
		return nil, ErrorUserNoneCap
	}
	if len(changeMap) != 0 {
		//融合重合部分
		for _, c := range capabilities {
			c.Score = c.Score + changeMap[c.Name]
		}
	}
	return capabilities, nil
}

// ChangeCapability 处理改变能力值逻辑
func ChangeCapability(userID int64, name string, change int16) (err error) {
	var (
		dScore int16
	)
	//从redis查询是否存在
	dScore, err = redis.GetCapScoreByName(userID, name)
	if err != nil { //不存在从mysql检查是否存在此能力
		if errors.Is(err, redis.ErrorNoneCap) {
			zap.L().Info("redis dont have capability",
				zap.Int64("userID", userID))
			//mysql检查是否存在
			exist, err2 := mysql.CheckUserCapExist(userID, name)
			if err2 != nil {
				return err2
			}
			if !exist {
				zap.L().Error("user capability not exist",
					zap.Int64("userID", userID),
					zap.String("capability_name", name))
				return ErrorCapNotExist
			}
			dScore = 0
			return scoreChange(userID, name, dScore, change)
		}
		return err
	}
	//存在改变数值
	return scoreChange(userID, name, dScore, change)
}

// DeleteCapByName 处理改变能力值逻辑
func DeleteCapByName(userID int64, name string) error {
	return nil
}

// 处理分数变化
func scoreChange(userID int64, name string, dScore, change int16) (err error) {
	return redis.StoreCapScore(userID, name, dScore+change)
}
