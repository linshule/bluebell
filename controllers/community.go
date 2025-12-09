package controllers

import (
	"net/http"
	"strconv"

	"bluebell/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, "服务器繁忙")
		return
	}
	c.JSON(http.StatusOK, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "无效的ID")
		return
	}

	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, http.StatusInternalServerError, "服务器繁忙")
		return
	}
	c.JSON(http.StatusOK, data)
}
