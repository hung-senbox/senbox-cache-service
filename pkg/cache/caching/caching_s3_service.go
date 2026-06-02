package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingS3Service interface {
	SetS3UrlByKey(ctx context.Context, key string, url string) error
	InvalidateS3UrlByKey(ctx context.Context, key string) error
}

type cachingS3Service struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingS3Service(cache *cache.RedisCache, defaultTTL int) CachingS3Service {
	return &cachingS3Service{cache: cache, defaultTTL: defaultTTL}
}

func (s *cachingS3Service) setByKey(ctx context.Context, key string, data string) error {
	if data == "" {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingS3Service) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingS3Service) SetS3UrlByKey(ctx context.Context, key string, url string) error {
	if key == "" || url == "" {
		return nil
	}
	return s.setByKey(ctx, keys.GetS3UrlByKeyCacheKey(key), url)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingS3Service) InvalidateS3UrlByKey(ctx context.Context, key string) error {
	if key == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetS3UrlByKeyCacheKey(key))
}
