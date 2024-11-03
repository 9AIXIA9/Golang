package redis

// redis key
// redis key注意使用命名空间的方式，方便查询和拆分
const (
	KeyPrefix  = "life:"
	KeyToken   = "token"
	KeyUserCap = "user:cap:"
)

// 给redis的key加前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
