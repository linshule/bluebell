package main

import (
	"bluebell/config"
	"bluebell/dao/mysql"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"fmt"

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

	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	r := routes.Setup()

	port := viper.GetString("app.port")

	zap.L().Info("服务器启动...", zap.String("port", port))

	r.Run(":" + port)
}
