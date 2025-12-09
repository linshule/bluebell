package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func SignUp(p *models.ParamSignUp) (err error) {
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}
