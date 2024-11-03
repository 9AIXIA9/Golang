package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost 用于创建帖子
func CreatePost(p *models.Post) (err error) {
	//生成post_id
	p.ID = snowflake.GenID()
	//2.保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	return redis.CreatePost(p.ID, p.CommunityID)

}

// GetPostByID 根据帖子id获取数据
func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	var (
		post      *models.Post
		user      *models.User
		community *models.CommunityDetail
	)
	data = new(models.ApiPostDetail)
	post, err = mysql.GetPostByID(pid)
	if err != nil {
		return
	}
	//根据作者id获取作者信息
	user, err = mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.Int64("author_id", post.AuthorID)
		return
	}
	//根据社区id获取社区信息
	community, err = mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.Int64("community_id", post.CommunityID)
		return
	}
	data.Post = post
	data.AuthorName = user.Username
	data.CommunityDetail = community
	return
}

// GetPostList 获取帖子数据
func GetPostList(pageNum, pageSize int64) (data []*models.ApiPostDetail, err error) {

	posts, err := mysql.GetPostList(pageNum, pageSize) //通过postID获取posts
	if err != nil {
		return nil, err
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		//根据作者id获取作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.Int64("author_id", post.AuthorID)
			continue
		}
		//根据社区id获取社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.Int64("community_id", post.CommunityID)
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// GetPostList2 获取帖子数据升级版
func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//去redis查询列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return 0 data")
		return
	}
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	//去mysql获取post详情
	posts, err := mysql.GetPostListByIDs(ids)

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for i, post := range posts {
		//根据作者id获取作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.Int64("author_id", post.AuthorID)
			continue
		}
		//根据社区id获取社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.Int64("community_id", post.CommunityID)
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[i],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, nil
}

// GetCommunityPostList 根据社区和顺序查询帖子
func GetCommunityPostList(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	//去redis查询列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return 0 data")
		return
	}
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	//去mysql获取post详情
	posts, err := mysql.GetPostListByIDs(ids)

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for i, post := range posts {
		//根据作者id获取作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.Int64("author_id", post.AuthorID)
			continue
		}
		//根据社区id获取社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.Int64("community_id", post.CommunityID)
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[i],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, nil
}
