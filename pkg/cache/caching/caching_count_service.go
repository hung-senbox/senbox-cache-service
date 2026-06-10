package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/cache/cached"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingCountService interface {
	SetMediaPortalCountByStudentID(ctx context.Context, studentID string, data interface{}) error
	InvalidateMediaPortalCountByStudentID(ctx context.Context, studentID string) error
}

type cachingCountService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingCountService(cache *cache.RedisCache, defaultTTL int) CachingCountService {
	return &cachingCountService{cache: cache, defaultTTL: defaultTTL}
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingCountService) SetMediaPortalCountByStudentID(ctx context.Context, studentID string, data interface{}) error {
	if studentID == "" || data == nil {
		return nil
	}
	return cached.SetCache(s.cache, ctx, keys.GetCountCacheKey(studentID), data)
}

func (s *cachingCountService) InvalidateMediaPortalCountByStudentID(ctx context.Context, studentID string) error {
	if studentID == "" {
		return nil
	}
	return cached.DeleteCache(s.cache, ctx, keys.GetCountCacheKey(studentID))
}
