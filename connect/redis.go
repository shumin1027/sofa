package connect

import (
	"github.com/go-redis/redis/v7"
)

// redis配置信息
type RedisConfig struct {
	Address  []string
	Password string
	DB       int
}

var redisClient redis.UniversalClient

// 初始化Redis客户端实例
func InitRedis(rc *RedisConfig) error {
	redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      rc.Address,
		Password:   rc.Password,
		DB:         rc.DB,
		MaxRetries: 2,
	})
	_, err := redisClient.Ping().Result()
	return err
}

// 获取Redis客户端实例
func RedisClient() redis.UniversalClient {
	return redisClient
}
