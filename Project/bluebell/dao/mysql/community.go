package mysql

import (
	"bluebell/models"
	"bluebell/myerrors"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

// GetCommunityList 获取数据库中社区数据
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `SELECT community_id,community_name from community`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn(myerrors.CommunityNoData.Error())
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 通过id获取数据库中的社区数据
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	//初始化变量
	community = new(models.CommunityDetail)
	sqlStr := `SELECT community_id,community_name,introduction,create_time FROM community where community.community_id = ?`
	if err = db.Get(community, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = myerrors.CommunityInvalidID
		}
	}
	return community, err
}
