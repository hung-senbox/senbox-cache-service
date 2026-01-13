package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedLocationService interface {
	GetLocationById(ctx context.Context, locationID uint) (map[string]interface{}, error)
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
// === HELPER METHODS ===
// ========================
func (c *cachedLocationService) getCache(
	ctx context.Context,
	cacheKey string,
) (map[string]interface{}, error) {
	if cacheKey == "" {
		return nil, nil
	}

	var result map[string]interface{}
	if err := c.cache.Get(ctx, cacheKey, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedLocationService) GetLocationById(ctx context.Context, locationID uint) (map[string]interface{}, error) {
	if locationID == 0 {
		return nil, nil
	}
	return c.getCache(ctx, keys.GetLocationByIdCacheKey(locationID))
}
