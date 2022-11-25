package initialize

import (
	"acc/server/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		return
	}
	global.GVA_LOG.Info("Redis 连接成功: " + pong)
	global.GVA_REDIS = client
}
