package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList 获取社区数据名单
func GetCommunityList() ([]*models.Community, error) {
	//查询数据库，查到所有community返回
	return mysql.GetCommunityList()
}

// GetCommunityDetailByID 通过id获取特定社区的数据
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
