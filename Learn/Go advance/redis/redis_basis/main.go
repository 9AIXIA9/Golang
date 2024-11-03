package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb *redis.Client

func initRDB() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	_, err = rdb.Ping(nil).Result()
	return err
}

// 常见操作
func redisExample() {
	err := rdb.Set(nil, "coke", 100, 0).Err()
	if err != nil {
		fmt.Printf("rdb get failed,err:%v", err)
		return
	}
	val, err := rdb.Get(nil, "coke").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("coke does not exist")
	} else if err != nil {
		fmt.Printf("get coke failed,err:%v", err)
		return
	} else {
		fmt.Println("coke", val)
	}
	fmt.Println(val)

	val2, err := rdb.Get(nil, "name").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed,err:%v", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

// 哈希表操作
func hgetDemo() {
	v, err := rdb.HGetAll(nil, "user").Result() //获取哈希表所有字段
	if err != nil {
		//redis.nil
		//或者其他错误
		fmt.Printf("hgetall failed,err:%v\n", err)
		return
	}
	fmt.Println(v)
	v2 := rdb.HMGet(nil, "user", "age", "name").Val() //拿到哈希表对应字段
	fmt.Println(v2)
	v3 := rdb.HGet(nil, "user", "age").Val()
	fmt.Println(v3)
}

// doDemo rdb.Do 方法使用示例
func doDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 直接执行命令获取错误
	err := rdb.Do(ctx, "set", "key", 10, "EX", 3600).Err()
	fmt.Println(err)

	// 执行命令获取结果
	val, err := rdb.Do(ctx, "get", "key").Result()
	fmt.Println(val, err)
}

// zsetDemo 操作zset示例
func zsetDemo() {
	// key
	zsetKey := "language_rank"
	// value
	// 注意：v8版本使用[]*redis.Z；此处为v9版本使用[]redis.Z
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// scanKeysDemo1 按前缀查找所有key示例
func scanKeysDemo1() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		// 将redis中所有以prefix:为前缀的key都扫描出来
		keys, cursor, err = rdb.Scan(ctx, cursor, "prefix:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
}

// scanKeysDemo2 按前缀扫描key示例
// 针对这种需要遍历大量key的场景，go-redis中提供了一个简化方法——Iterator
func scanKeysDemo2() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// 按前缀扫描key
	iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

func main() {

	if err := initRDB(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connect redis success...")
	//程序退出时释放相关资源
	defer rdb.Close()
	//redisExample()
	//hgetDemo()
	zsetDemo()
	scanKeysDemo1()
}
