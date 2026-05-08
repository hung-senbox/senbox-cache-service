package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedColortimeService interface {
	GetColorTimeSessionID(ctx context.Context, colorTimeSessionID string) (map[string]interface{}, error)
}

type cachedColortimeService struct {
	cache *cache.RedisCache
}

func NewCachedColortimeService(cache *cache.RedisCache) CachedColortimeService {
	return &cachedColortimeService{cache: cache}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedColortimeService) GetColorTimeSessionID(ctx context.Context, colorTimeSessionID string) (map[string]interface{}, error) {
	if colorTimeSessionID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetColorTimeSessionIDCacheKey(colorTimeSessionID))
}
