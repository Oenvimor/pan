package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log/slog"
)

var redisPool *redis.Pool

func InitRedisPool() {
	redisPool = &redis.Pool{
		MaxIdle:     viper.GetInt("redis.max_idle"),
		MaxActive:   viper.GetInt("redis.max_active"),
		IdleTimeout: viper.GetDuration("redis.idle_timeout"),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")))
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
