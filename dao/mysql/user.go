package mysql

import (
	"bluebell/models"
	"errors"
)

func CheckUserExist(username string) (err error) {
	var count int64
	err = DB.Model(&models.User{}).Where("username=?", username).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

func InsertUser(user *models.User) (err error) {
	err = DB.Create(user).Error
	return
}
