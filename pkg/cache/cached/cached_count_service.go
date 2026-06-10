package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedCountService interface {
	GetMediaPortalCountByStudentID(ctx context.Context, studentID string) (map[string]interface{}, error)
}

type cachedCountService struct {
	cache *cache.RedisCache
}

func NewCachedCountService(cache *cache.RedisCache) CachedCountService {
	return &cachedCountService{cache: cache}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedCountService) GetMediaPortalCountByStudentID(ctx context.Context, studentID string) (map[string]interface{}, error) {
	if studentID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetCountCacheKey(studentID))
}
