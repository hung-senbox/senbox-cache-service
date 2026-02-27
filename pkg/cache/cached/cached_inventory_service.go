package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedInventoryService interface {
	GetShelfSlotById(ctx context.Context, shelfSlotID string) (map[string]interface{}, error)
}

type cachedInventoryService struct {
	cache *cache.RedisCache
}

func NewCachedInventoryService(cache *cache.RedisCache) CachedInventoryService {
	return &cachedInventoryService{cache: cache}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedInventoryService) GetShelfSlotById(ctx context.Context, shelfSlotID string) (map[string]interface{}, error) {
	if shelfSlotID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetShelfSlotByIdCacheKey(shelfSlotID))
}
