package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()
	p.Status = 1
	err = mysql.CreatePost(p)
	if err != nil {
		zap.L().Error("mysql.CreatePost failed", zap.Error(err))
	}
	return
}

// GetPostDetail 获取帖子详情
func GetPostDetail(pid int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById failed", zap.Error(err))
		return
	}

	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID failed", zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName: user.Username,
		Post:       post,
		Community:  community,
	}
	return
}
