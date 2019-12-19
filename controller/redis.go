package controller

import (
	model "RedisMonitor/model"
	"fmt"
	"gopkg.in/redis.v4"
	"log"
)

var (
	ConfigData *model.Config
)

// 取得 Redis 連線
func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", ConfigData.Redis.Url, ConfigData.Redis.Port),
		Password: ConfigData.Redis.Password,
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatalf("創建 Redis Client 時發生錯誤，錯誤訊息為 %s", err)
	}

	return client
}
