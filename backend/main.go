package main

import (
	"bbs/dao/mysql"
	"bbs/dao/redis"
	"bbs/logger"
	"bbs/routes"
	"bbs/settings"
	"github.com/spf13/viper"
)

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		panic(err)
	}
	// 日志配置
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		panic(err)
	}
	// 连接数据库
	if err := mysql.Init(); err != nil {
		panic(err)
	}
	// 连接redis
	if err := redis.Init(); err != nil {
		panic(err)
	}
	// 初始化路由
	r := routes.Init()
	err := r.Run("0.0.0.0:" + viper.GetString("app.port"))
	if err != nil {
		return
	}
}
