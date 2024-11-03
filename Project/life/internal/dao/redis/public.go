package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

// 数据库
var (
	rdb *redis.Client
)

const (
	capStoreTime = time.Hour
	onceNum      = 20
)

// 错误
var (
	errorRedis = errors.New("redis init failed")
)
