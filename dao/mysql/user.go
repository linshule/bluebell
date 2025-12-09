package mysql

import (
	"bluebell/models"
	"bluebell/pkg/encrypt"
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

func Login(user *models.User) (err error) {
	originPassword := user.Password

	// SQL: SELECT * FROM users WHERE username = ? LIMIT 1
	dbErr := DB.Where("username = ?", user.Username).First(user).Error
	if dbErr != nil {
		return errors.New("用户不存在")
	}

	err = encrypt.ComparePasswords(user.Password, originPassword)
	return
}
