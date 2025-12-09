package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/encrypt"
	"bluebell/pkg/jwt"
)

func SignUp(p *models.ParamSignUp) (err error) {
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	encryptedPassword, err := encrypt.EncryptPassword(p.Password)
	if err != nil {
		return err
	}
	user := &models.User{
		Username: p.Username,
		Password: encryptedPassword,
	}
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		return "", err
	}

	return jwt.GenToken(int64(user.ID), user.Username)
}
