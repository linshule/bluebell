package models

import "time"

type Post struct {
	ID          int64     `json:"id" gorm:"column:post_id;primaryKey"`
	AuthorID    int64     `json:"author_id" gorm:"column:author_id"`
	CommunityID int64     `json:"community_id" gorm:"column:community_id" binding:"required"`
	Status      int32     `json:"status" gorm:"column:status"`
	Title       string    `json:"title" gorm:"column:title" binding:"required"`
	Content     string    `json:"content" gorm:"column:content" binding:"required"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
}

func (Post) TableName() string {
	return "post"
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Post
	*Community `json:"community"`
}

type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
