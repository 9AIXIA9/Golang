package redis

import (
	"fmt"
	"life/internal/settings"

	"github.com/go-redis/redis"
)

// Init 初始化redis连接
func Init() (err error) {
	config := settings.Config.RedisConf
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port),
		Password: config.Password, // 密码
		DB:       config.Database, // 数据库
		PoolSize: config.PoolSize, // 连接池大小
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		err = fmt.Errorf("%w,connect redis failed,err:%w", errorRedis, err)
	}
	return
}

func Close() error {
	return rdb.Close()
}
