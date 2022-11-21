package redis

import (
	"acc/config"
	"acc/lib/logger"
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func Setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Host + ":" + config.RedisConfig.Port,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})

	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("Redis 连接失败: %v", err)
	}
	logger.Debug("Redis 连接成功: " + pong)
}

func Rdb() *redis.Client {
	return rdb
}
