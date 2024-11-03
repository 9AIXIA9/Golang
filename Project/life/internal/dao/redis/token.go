package redis

import (
	"errors"
	"fmt"
	"life/internal/settings"
	"time"

	"github.com/go-redis/redis"

	"go.uber.org/zap"
)

// GetTokenByUserID 通过userID获取token
func GetTokenByUserID(userID int64) (token string, err error) {
	userIDStr := fmt.Sprintf("%d", userID)
	token, err = rdb.HGet(getRedisKey(KeyToken), userIDStr).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			zap.L().Error(ErrorNoneToken.Error(),
				zap.Int64("userID", userID))
			return "", ErrorNoneToken
		}
		zap.L().Error("redis: get user token failed",
			zap.Int64("userID", userID),
			zap.Error(err))
		return "", err
	}
	return token, nil
}

// StoreToken 存储token到redis
func StoreToken(userID int64, token string) (err error) {
	tokenExistTime := time.Duration(settings.Config.TokenDuration) * time.Hour
	userIDStr := fmt.Sprintf("%d", userID)
	// 使用Hash存储userID对应的token
	err = rdb.HSet(getRedisKey(KeyToken), userIDStr, token).Err()
	if err != nil {
		zap.L().Error("redis: store token failed",
			zap.Int64("userID", userID),
			zap.Error(err))
		return err
	}
	err = rdb.Expire(getRedisKey(KeyToken), tokenExistTime).Err()

	if err != nil {
		zap.L().Error("redis expire token failed",
			zap.Int64("userID", userID),
			zap.Error(err))
		return err
	}
	return
}
