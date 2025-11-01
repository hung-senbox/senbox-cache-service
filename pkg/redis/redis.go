package redis

import (
	"context"
	"fmt"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	goredis "github.com/redis/go-redis/v9"
)

func InitRedisCache(host, port, password string, db int) (*cache.RedisCache, error) {
	addr := fmt.Sprintf("%s:%s", host, port)

	client := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Kiểm tra kết nối Redis
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("cannot connect to redis (%s): %w", addr, err)
	}

	return cache.NewRedisCache(client), nil
}
