package redis

import (
	"bluebell/models"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// CreatePost 在redis里面存储帖子时间,分数,社区
func CreatePost(postID, communityID int64) (err error) {
	postIDStr := fmt.Sprintf("%d", postID)

	err = rdb.ZAdd(getRedisKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postIDStr,
	}).Err()

	err = rdb.ZAdd(getRedisKey(KeyPostScore), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postIDStr,
	}).Err()

	err = rdb.SAdd(getRedisKey(KeyCommunityPostPrefix+strconv.Itoa(int(communityID))), postIDStr).Err()
	return
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis获取id

	//根据用户选择查询的key
	key := getRedisKey(KeyPostTime)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScore)
	}
	return getIDSFromKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedPrefix + id)
	//	//查询key分数是1的数量->赞成票数量
	//	v := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	//使用pipeline一次性发送多条命令，节省RTT
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}

// GetCommunityPostIDsInOrder 按社区和顺序查询帖子ID
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	// 1. 确定排序的key
	orderKey := getRedisKey(KeyPostTime) // 默认按时间排序
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScore) // 按分数排序
	}

	// 2. 构造社区的key
	communityKey := getRedisKey(KeyCommunityPostPrefix + strconv.Itoa(int(p.CommunityID)))

	// 3. 构造缓存key，格式：community_post_order:{communityID}:{order}
	cacheKey := fmt.Sprintf("%s%d:%s", KeyCommunityPostPrefix, p.CommunityID, p.Order)

	// 4. 判断缓存是否存在
	if rdb.Exists(cacheKey).Val() < 1 {
		pipeline := rdb.Pipeline()

		// 5. 使用 ZInterStore 做交集运算
		// 注意：这里需要使用 destination, keys... 的格式
		pipeline.ZInterStore(cacheKey,
			redis.ZStore{
				Aggregate: "MAX",
			},
			communityKey, orderKey) // 直接传入需要交集的key

		// 6. 设置缓存过期时间
		pipeline.Expire(cacheKey, 60*time.Second)

		// 7. 执行 pipeline
		_, err := pipeline.Exec()
		if err != nil {
			return nil, fmt.Errorf("执行社区帖子列表计算失败: %w", err)
		}
	}

	// 8. 从计算结果中获取指定页的数据
	return getIDSFromKey(cacheKey, p.Page, p.Size)
}

// 辅助函数：根据ids获取指定范围的元素
func getIDSFromKey(key string, page, size int64) ([]string, error) {
	//确定起始点
	start := (page - 1) * size
	end := start + size - 1
	//获取列表
	return rdb.ZRevRange(key, start, end).Result()
}
