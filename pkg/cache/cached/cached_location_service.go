package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedLocationService interface {
	GetLocationById(ctx context.Context, locationID string) (map[string]interface{}, error)
}

type cachedLocationService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachedLocationGateway(cache *cache.RedisCache) CachedLocationService {
	return &cachedLocationService{
		cache: cache,
	}
}


// ========================
// === GET CACHE ===
// ========================

func (c *cachedLocationService) GetLocationById(ctx context.Context, locationID string) (map[string]interface{}, error) {
	if locationID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetLocationByIdCacheKey(locationID))
}
