package models

import "time"

type Community struct {
	ID   int64  `json:"id,string" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID   int64  `json:"id,string" db:"community_id"`
	Name string `json:"name" db:"community_name"`
	//omitempty的意思是如果字段为空则不展示
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
