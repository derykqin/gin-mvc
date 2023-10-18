package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func InitRedis() {
	ctx := context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "12345678",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err == redis.Nil {
		log.Fatal("Redis异常")
	} else if err != nil {
		log.Fatal("ping redis 失败")
	} else {
		log.Info("Redis连接成功")
	}
}
