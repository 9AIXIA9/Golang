package redis

import (
	"errors"
	"fmt"
	"life/pkg/models"
	"strconv"
	"strings"

	"github.com/go-redis/redis"

	"go.uber.org/zap"
)

// StoreCapScore 存储能力值
func StoreCapScore(userID int64, capName string, capScore int16) error {
	key := getRedisKey(KeyUserCap + fmt.Sprintf("%d", userID))

	pipe := rdb.Pipeline()
	defer func() {
		if err := pipe.Close(); err != nil {
			zap.L().Error("redis pipeline close failed")
		}
	}()

	// 使用Pipeline批量执行命令
	pipe.HSet(key, capName, capScore)
	pipe.Expire(key, capStoreTime)

	_, err := pipe.Exec()
	if err != nil {
		zap.L().Error("store user capability failed",
			zap.Int64("user_id", userID),
			zap.String("capability", capName),
			zap.Error(err))
		return fmt.Errorf("store capability score: %w", err)
	}

	return nil
}

// GetCapScoreByName 获取用户能力值通过名字
func GetCapScoreByName(userID int64, capName string) (int16, error) {
	key := getRedisKey(KeyUserCap + fmt.Sprintf("%d", userID))

	score, err := rdb.HGet(key, capName).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			zap.L().Info("capability not found",
				zap.Int64("user_id", userID),
				zap.String("cap_name", capName))
			return 0, ErrorNoneCap
		}
		zap.L().Error("get user capability score failed",
			zap.Int64("user_id", userID),
			zap.Error(err))
		return 0, fmt.Errorf("get capability score: %w", err)
	}

	s, err := strconv.ParseInt(score, 10, 16)
	if err != nil {
		zap.L().Error("parse capability score failed",
			zap.Int64("user_id", userID),
			zap.String("score", score),
			zap.Error(err))
		return 0, fmt.Errorf("parse capability score: %w", err)
	}

	return int16(s), nil
}

// GetAllCapScore 获取用户所有能力值
func GetAllCapScore(userID int64) (map[string]int16, error) {
	key := getRedisKey(KeyUserCap + fmt.Sprintf("%d", userID))
	data, err := rdb.HGetAll(key).Result()
	if err != nil {
		zap.L().Error("get all user capabilities failed",
			zap.Int64("user_id", userID),
			zap.Error(err))
		return nil, fmt.Errorf("get all capabilities: %w", err)
	}
	if len(data) == 0 {
		zap.L().Info("no capabilities found",
			zap.Int64("user_id", userID))
		return nil, ErrorNoneCap
	}
	//把data转化为[string]int16
	caps := make(map[string]int16)
	for name, score := range data {
		s, err := strconv.ParseInt(score, 10, 16)
		if err != nil {
			zap.L().Error("parse capability score failed",
				zap.Int64("user_id", userID),
				zap.String("cap_name", name),
				zap.String("score", score),
				zap.Error(err))
			return nil, fmt.Errorf("parse capability score: %w", err)
		}
		caps[name] = int16(s)
	}

	return caps, nil
}

// DeleteCapByID 删除用户所有能力值
func DeleteCapByID(userID int64) error {
	k := getRedisKey(KeyUserCap + fmt.Sprintf("%d", userID))
	if err := rdb.Del(k).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrorNoneCap
		}
		zap.L().Error("delete redis capability failed", zap.Int64("userID", userID))
		return err
	}
	return nil
}

// QueryAllUser 查询所有用户在redis中的能力值
func QueryAllUser() ([]models.UserCapability, error) {
	// 使用SCAN命令遍历所有用户键
	var cursor uint64 // Redis scan 游标
	var users []models.UserCapability

	for {
		// 每次扫描100个键
		keys, nextCursor, err := rdb.Scan(cursor, getRedisKey(KeyUserCap)+"*", onceNum).Result()
		if err != nil {
			zap.L().Error("scan redis keys failed",
				zap.Error(err))
			return nil, fmt.Errorf("scan redis keys: %w", err)
		}

		// 使用Pipeline批量获取用户数据
		pipe := rdb.Pipeline()

		// 为每个键创建HGETALL命令
		cmds := make(map[string]*redis.StringStringMapCmd, len(keys))
		//HGETALL 返回一个 hash 表的所有字段和值
		//结果是一个 map[string]string，表示字段名和对应的值
		for _, key := range keys {
			cmds[key] = pipe.HGetAll(key)
		}

		// 执行Pipeline
		_, err = pipe.Exec()

		// 立即关闭 pipeline，而不是使用 defer
		if closeErr := pipe.Close(); closeErr != nil {
			zap.L().Error("redis pipeline close failed",
				zap.Error(closeErr))
		}

		if err != nil {
			zap.L().Error("execute pipeline failed",
				zap.Error(err))
			return nil, fmt.Errorf("execute pipeline: %w", err)
		}

		// 处理结果
		for key, cmd := range cmds {
			capabilities, err := cmd.Result()
			if err != nil {
				zap.L().Error("get user capabilities failed",
					zap.String("key", key),
					zap.Error(err))
				continue
			}

			// 从键中提取用户ID
			userIDStr := strings.TrimPrefix(key, getRedisKey(KeyUserCap))
			userID, err := strconv.ParseInt(userIDStr, 10, 64)
			if err != nil {
				zap.L().Error("parse user id failed",
					zap.String("user_id_str", userIDStr),
					zap.Error(err))
				continue
			}

			// 如果能力值不为空，添加到结果中
			if len(capabilities) > 0 {
				users = append(users, models.UserCapability{
					UserID:       userID,
					Capabilities: capabilities,
				})
			}
		}

		// 检查是否完成所有键的扫描
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return users, nil
}
