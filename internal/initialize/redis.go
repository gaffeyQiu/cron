package initialize

import (
	"context"
	"fmt"
	"tsf-cron/config"
	"tsf-cron/pkg/core/log"

	"github.com/go-redis/redis/v8"
)

func GetRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GetString("Redis.Ip"), config.GetInt("Redis.Port")),
		Password: config.GetString("Redis.Passwd"),
		DB:       0,
	})

	if err := rdb.Info(context.Background()).Err(); err != nil {
		log.Fatalf("Connect Redis Fail: %s", err)
	}

	return rdb
}
