package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingColortimeService interface {
	SetColorTimeSessionID(ctx context.Context, colorTimeSessionID string, data interface{}) error
	InvalidateColorTimeSessionID(ctx context.Context, colorTimeSessionID string) error
}

type cachingColortimeService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingColortimeService(cache *cache.RedisCache, defaultTTL int) CachingColortimeService {
	return &cachingColortimeService{cache: cache, defaultTTL: defaultTTL}
}

func (s *cachingColortimeService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingColortimeService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingColortimeService) SetColorTimeSessionID(ctx context.Context, colorTimeSessionID string, data interface{}) error {
	if colorTimeSessionID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.GetColorTimeSessionIDCacheKey(colorTimeSessionID), data)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingColortimeService) InvalidateColorTimeSessionID(ctx context.Context, colorTimeSessionID string) error {
	if colorTimeSessionID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetColorTimeSessionIDCacheKey(colorTimeSessionID))
}
