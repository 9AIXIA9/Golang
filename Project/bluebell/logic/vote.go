package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/myerrors"
	"bluebell/settings"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

//本系统使用简化版算法--阮一峰有更强的
//投一票就加432分 86400/200-> 200张赞成票可以给帖子续一天

/*
投票的几种情况:
没投过 ,  投过
赞成  不投  反对
  1   0   -1
投票限制:
每个帖子只能在发表后一星期内投票，超过不让投票
到期将redis库中数据转移到mysql，并释放redis内存
*/

// PostVote 用于处理帖子投票功能
func PostVote(p *models.ParamVoteData, userID int64) error {
	// 1. 判断是否在投票期限
	if err := CheckPostVoteTime(p.PostID); err != nil {
		return errors.Wrap(err, "checking post vote time")
	}

	// 2. 更新数据
	if err := redis.PostVote(p, userID); err != nil {
		return errors.Wrap(err, "updating vote in Redis")
	}

	// 3. 返回
	return nil
}

// CheckPostVoteTime 检查帖子是否在投票期限内
func CheckPostVoteTime(postID int64) error {
	createTime, err := redis.GetPostTime(postID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("getting create time for post %d", postID))
	}

	maxDuration := getPostVoteDuration()
	if time.Since(createTime) > maxDuration {
		return myerrors.PostVoteOverdue
	}

	return nil
}

func getPostVoteDuration() time.Duration {
	return time.Hour * time.Duration(settings.Config.Set.PostVoteDuration)
}
