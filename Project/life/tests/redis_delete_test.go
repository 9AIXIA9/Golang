package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func TestDeleteRedis(t *testing.T) {
	InitRedis()
	if err := DeleteCapByName(); err != nil {
		log.Print(err)
	}
}

// InitRedis 初始化
func InitRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("连接Redis出现错误：" + err.Error())
	} else {
		fmt.Println("连接Redis成功:" + pong)
	}
	return rdb
}

// DeleteCapByName 删除用户能力值 直接删除所有
func DeleteCapByName() error {
	k := "life:user:4374229646577664"
	rdb.Del(k)
	return nil
}
