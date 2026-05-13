package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingTermService interface {
	// Set term cache
	SetTermCache(ctx context.Context, termID string, data interface{}) error
	// Invalidate term cache
	InvalidateTermCache(ctx context.Context, termID string) error
}

type cachingTermService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingTermService(redisCache *cache.RedisCache, defaultTTL int) CachingTermService {
	return &cachingTermService{cache: redisCache, defaultTTL: defaultTTL}
}

func (s *cachingTermService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingTermService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingTermService) SetTermCache(ctx context.Context, termID string, data interface{}) error {
	if termID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.GetTermCacheKey(termID), data)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingTermService) InvalidateTermCache(ctx context.Context, termID string) error {
	if termID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetTermCacheKey(termID))
}
