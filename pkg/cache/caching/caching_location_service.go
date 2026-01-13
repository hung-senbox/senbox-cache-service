package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingLocationService interface {
	SetLocationById(ctx context.Context, locationID string, data interface{}) error
	InvalidateLocationById(ctx context.Context, locationID string) error
}

type cachingLocationService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingLocationService(redisCache *cache.RedisCache, defaultTTL int) CachingLocationService {
	return &cachingLocationService{
		cache:      redisCache,
		defaultTTL: defaultTTL,
	}
}

func (s *cachingLocationService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingLocationService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingLocationService) SetLocationById(ctx context.Context, locationID string, data interface{}) error {
	if locationID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.GetLocationByIdCacheKey(locationID), data)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingLocationService) InvalidateLocationById(ctx context.Context, locationID string) error {
	if locationID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetLocationByIdCacheKey(locationID))
}
