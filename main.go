package main

import (
	"bluebell/config"
	"bluebell/dao/mysql"
	"bluebell/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {

	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := logger.Init(); err != nil {
		panic(err)
	}

	defer zap.L().Sync()

	zap.L().Info("配置和日志已加载成功！")
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	zap.L().Info("MySQL连接成功！")
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		appName := viper.GetString("app.name")

		zap.L().Debug("接收到hello请求")

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello," + appName,
			"version": viper.GetString("app.version"),
		})
	})

	port := viper.GetString("app.port")

	zap.L().Info("服务器启动...", zap.String("port", port))

	r.Run(":" + port)
}
