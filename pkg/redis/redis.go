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

	rc := cfg.Database.RedisCache

	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", rc.Host, rc.Port),
		Password: rc.Password,
		DB:       rc.DB,
	})

	// Check connect Redis
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("cannot connect to redis: %w", err)
	}

	return cache.NewRedisCache(client), nil
}
