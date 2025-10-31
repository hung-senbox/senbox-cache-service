package redis

import (
	"context"
	"fmt"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/config"
	goredis "github.com/redis/go-redis/v9"
)

func InitRedisCacheFromFile(path string) (*cache.RedisCache, error) {
	cfg, err := config.LoadConfig(path)
	if err != nil {
		return nil, err
	}

	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Test kết nối
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("cannot connect to redis: %w", err)
	}

	return cache.NewRedisCache(client), nil
}
