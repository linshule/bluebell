package mysql

import (
	"bluebell/models"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	err = DB.Create(p).Error
	return
}

// GetPostById 根据ID查询帖子详情
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	err = DB.Where("post_id = ?", pid).First(post).Error
	return
}
