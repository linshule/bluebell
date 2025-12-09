package mysql

import (
	"bluebell/models"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetCommunityList 查询社区列表
func GetCommunityList() (data []*models.Community, err error) {
	if err = DB.Find(&data).Error; err != nil {
		zap.L().Error("there is no community in db", zap.Error(err))
		return nil, err
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.Community, err error) {
	community = new(models.Community)
	if err = DB.Where("community_id = ?", id).First(community).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("社区不存在")
		}
	}
	return
}
