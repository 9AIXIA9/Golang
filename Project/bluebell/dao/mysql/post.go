package mysql

import (
	"bluebell/models"
	"bluebell/myerrors"
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
)

// CreatePost 执行post入库
func CreatePost(p *models.Post) (err error) {
	sqlStr := `INSERT INTO post (post_id,title,content,author_id,community_id)  VALUES(?,?,?,?,?)`

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostByID 根据id获取单个post数据
func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `SELECT post_id,title,content,author_id,community_id,create_time FROM post WHERE post_id = ?`
	err = db.Get(post, sqlStr, pid)
	if errors.Is(err, sql.ErrNoRows) {
		err = myerrors.PostNotExist
	}
	return
}

// GetPostList 按照创建时间获取post数据
func GetPostList(pageNum, pageSize int64) (posts []*models.Post, err error) {
	sqlStr := `SELECT post_id,title,content,author_id,community_id,create_time
	FROM post
	ORDER BY create_time
	DESC 
	LIMIT ?,?
	`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (pageNum-1)*pageSize, pageSize)
	if errors.Is(err, sql.ErrNoRows) {
		err = myerrors.PostNotExist
	}
	return
}

func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `SELECT post_id, title, content, author_id, community_id, status, create_time
	FROM post
	WHERE post_id IN (?)
	ORDER BY FIND_IN_SET(post_id,?) 
`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
