package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Community struct {
	ID           int64     `json:"id" gorm:"column:community_id"`
	Name         string    `json:"name" gorm:"column:community_name"`
	Introduction string    `json:"introduction" gorm:"column:introduction"`
	CreatedAt    time.Time `json:"create_time" gorm:"column:create_time"`
	UpdatedAt    time.Time `json:"update_time" gorm:"column:update_time"`
}

func (Community) TableName() string {
	return "community"
}
