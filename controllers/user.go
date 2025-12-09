package controllers

import (
	"net/http"

	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg":   "请求参数有误",
			"error": err.Error(),
		})
		return
	}

	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg":   "注册失败",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "参数无效"})
		return
	}

	token, err := logic.Login(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "登录失败",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": token,
	})
}
