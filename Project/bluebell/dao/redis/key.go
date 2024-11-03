package redis

// redis key
// redis key注意使用命名空间的方式，方便查询和拆分
const (
	KeyPrefix              = "bluebell:"
	KeyToken               = "user:token"
	KeyPostTime            = "post:time"   //zset;帖子及发帖时间
	KeyPostScore           = "post:score"  //zset;帖子及投票分数
	KeyPostVotedPrefix     = "post:voted:" //zset;参数是postID;用户及投票类型
	KeyCommunityPostPrefix = "community:"  //set保存每个分区的帖子ID
)

// 给redis的key加前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
