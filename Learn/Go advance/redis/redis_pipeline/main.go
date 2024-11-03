package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// PIPELINE Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能
// 区别于一个接一个地执行100个命令
// 你可以将这些命令放入 pipeline 中
// 然后使用1次读写操作像执行单个命令一样执行它们
// 这样做的好处是节省了执行命令的网络往返时间（RTT）
func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	//方法一
	pipe := rdb.Pipeline()
	var ctx context.Context
	ctx = nil
	incr := pipe.Incr(ctx, "pipeline_counter")
	//incr用于给key值加1
	pipe.Expire(ctx, "pipeline_counter", time.Hour)
	//设置过期时间

	cmds, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	//cmds用于存储执行结果
	fmt.Println(cmds)
	//方法二
	// 在执行pipe.Exec之后才能获取到结果
	fmt.Println(incr.Val())

	cmds, err = rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipelined_counter")
		pipe.Expire(ctx, "pipelined_counter", time.Hour)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// 在pipeline执行后获取到结果
	fmt.Println(incr.Val())

}
