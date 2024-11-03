package redis

import (
	"bluebell/settings"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// StoreToken 用于存储用户的token值到redis，确保每个用户只有一个token
func StoreToken(userID int64, token string) error {
	tokenExistTime := time.Duration(settings.Config.TokenDuration) * time.Hour
	userIDStr := fmt.Sprintf("%d", userID)

	// 使用Hash存储userID对应的token
	err := rdb.HSet(getRedisKey(KeyToken), userIDStr, token).Err()
	if err != nil {
		return fmt.Errorf("failed to store token: %w", err)
	}
	err = rdb.Expire(getRedisKey(KeyToken), tokenExistTime).Err()
	if err != nil {
		return fmt.Errorf("failed to set expiration: %w", err)
	}

	return nil
}

// GetTokenByUserID 通过userID获取redis中的token
func GetTokenByUserID(userID int64) (string, error) {
	userIDStr := fmt.Sprintf("%d", userID)
	token, err := rdb.HGet(getRedisKey(KeyToken), userIDStr).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", fmt.Errorf("token not found for userID %d", userID)
		}
		return "", fmt.Errorf("failed to get token for userID %d: %w", userID, err)
	}

	return token, nil
}
