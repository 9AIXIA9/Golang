package models

import "time"

//内存对齐

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id,string" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status" `
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情信息
type ApiPostDetail struct {
	AuthorName      string           `json:"author_name"`
	VoteNum         int64            `json:"vote_num"`
	Post            *Post            `json:"post"`      //嵌入帖子信息
	CommunityDetail *CommunityDetail `json:"community"` //嵌入社区信息
}
