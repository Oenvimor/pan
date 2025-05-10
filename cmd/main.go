package main

import (
	"log/slog"
	"pan/common/logger"
	"pan/config"
	db "pan/dao"
	"pan/router"
)

func main() {
	// 初始化日志
	logger.InitLogger("./log", slog.LevelDebug)
	// 加载配置文件
	config.InitConfig()
	// 连接数据库
	db.GetDB()
	defer db.Close()
	// 设置路由
	router.SetUpRouter()
}
