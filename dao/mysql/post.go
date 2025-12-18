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

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, size)

	offset := (page - 1) * size

	// GORM 查询：
	// Order("create_time desc"): 按创建时间倒序（新的在前）
	// Limit: 取多少条
	// Offset: 跳过多少条
	err = DB.Order("create_time desc").
		Limit(int(size)).
		Offset(int(offset)).
		Find(&posts).Error
	return
}
