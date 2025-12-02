package main

import (
	"bluebell/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	if err := config.Init(); err != nil {
		fmt.Printf("配置加载失败：%v\n", err)
		return
	}
	fmt.Printf("配置加载成功！\n")

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		appName := viper.GetString("app.name")

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello," + appName,
			"version": viper.GetString("app.version"),
		})
	})

	port := viper.GetString("app.port")
	r.Run(":" + port)
}
