package controllers

import (
	"net/http"
	"strconv"

	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		ResponseError(c, http.StatusUnauthorized, "请先登录")
		return
	}
	p.AuthorID = userID.(int64)
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, "创建帖子失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "无效的帖子ID")
		return
	}

	data, err := logic.GetPostDetail(pid)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "查询失败")
		return
	}
	c.JSON(http.StatusOK, data)
}
