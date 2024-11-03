package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" //初始化mysql驱动
)

var rdb *redis.Client

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port),
		Password: config.Password, // 密码
		DB:       config.Database, // 数据库
		PoolSize: config.PoolSize, // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
