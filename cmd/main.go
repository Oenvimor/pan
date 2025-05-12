package main

import (
	"log/slog"
	rdb "pan/cache"
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
	// 连接 MySQL
	db.InitDB()
	// 建立 Redis 连接池
	rdb.InitRedisPool()
	// 设置路由
	router.SetUpRouter()
}
