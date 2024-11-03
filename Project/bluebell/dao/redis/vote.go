package redis

import (
	"bluebell/models"
	"bluebell/myerrors"
	"bluebell/settings"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	voteScore float64 //每票分数
)

const (
	VoteNotChange = 0
	VoteReset     = 0
)

// PostVote 存储帖子投票数据
func PostVote(p *models.ParamVoteData, userID int64) (err error) {
	userIDStr := fmt.Sprintf("%d", userID)
	postIDStr := fmt.Sprintf("%d", p.PostID)
	value := float64(p.Direction)
	//查询当前用户投票记录
	oValue := rdb.ZScore(getRedisKey(KeyPostVotedPrefix+postIDStr), userIDStr).Val()
	// +1  0  -1

	//post分数变化
	scoreChange := (value - oValue) * getVoteScore()

	//检查投票是否发生变化
	if scoreChange == VoteNotChange {
		return myerrors.PostVoteNotChange
	}

	_, err = rdb.ZIncrBy(getRedisKey(KeyPostScore), scoreChange, postIDStr).Result()
	if err != nil {
		return errors.Join(myerrors.PostVote, err)
	}

	//记录用户投票
	if value == VoteReset {
		err = rdb.ZRem(getRedisKey(KeyPostVotedPrefix+postIDStr), userIDStr).Err()

	} else {
		err = rdb.ZAdd(getRedisKey(KeyPostVotedPrefix+postIDStr), redis.Z{
			Score:  value,
			Member: userIDStr,
		}).Err()
	}
	if err != nil {
		err = errors.Join(myerrors.PostVote, err)
	}
	return
}

func getVoteScore() float64 {
	if voteScore == 0 {
		voteScore = float64(settings.Config.Set.PostPerVoteScore)
	}
	return voteScore
}

func GetPostTime(postID int64) (createTime time.Time, err error) {
	postIDStr := fmt.Sprintf("%d", postID)
	var createTimeFloat float64
	createTimeFloat, err = rdb.ZScore(getRedisKey(KeyPostTime), postIDStr).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return time.Time{}, myerrors.PostTimeNotFound
		}
		return time.Time{}, err
	}
	createTime = time.Unix(int64(createTimeFloat), 0)
	return
}
