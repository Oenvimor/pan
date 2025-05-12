package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log/slog"
)

var redisPool *redis.Pool

func InitRedisPool() {
	redisPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				slog.Error("redis connect failed", "err", err)
				return nil, err
			}
			return c, nil
		},
	}
	TestRedisConnection()
}

func TestRedisConnection() {
	conn := redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", "key", "测试通过")
	if err != nil {
		slog.Error("Redis Test SET Failed", "err", err)
		return
	}
	value, err := redis.String(conn.Do("GET", "key"))
	if err != nil {
		slog.Error("Redis Test GET Failed", "err", err)
		return
	}
	slog.Info(fmt.Sprintf("Redis %s", value))
}
