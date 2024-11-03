package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb *redis.Client

func initRDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
}
func trans() {
	// TxPipeline demo
	var ctx context.Context
	ctx = nil
	pipe := rdb.TxPipeline()
	incr := pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)

	// TxPipelined demo
	var incr2 *redis.IntCmd
	_, err = rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		incr2 = pipe.Incr(ctx, "tx_pipeline_counter")
		pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
		return nil
	})
	fmt.Println(incr2.Val(), err)

}

// watchDemo 在key值不变的情况下将其值+1
func watchDemo(ctx context.Context, key string) error {
	return rdb.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时5秒
		// 5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(5 * time.Second)
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
}

func main() {
	initRDB()

}
