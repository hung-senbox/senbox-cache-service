package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedTermService interface {
	GetTermCache(ctx context.Context, termID string) (interface{}, error)
	GetCurrentTermByOrganizationIDCache(ctx context.Context, organizationID string) (interface{}, error)
}

type cachedTermService struct {
	cache *cache.RedisCache
}

func NewCachedTermService(cache *cache.RedisCache) CachedTermService {
	return &cachedTermService{cache: cache}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedTermService) GetTermCache(ctx context.Context, termID string) (interface{}, error) {
	if termID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetTermCacheKey(termID))
}

func (c *cachedTermService) GetCurrentTermByOrganizationIDCache(ctx context.Context, organizationID string) (interface{}, error) {
	if organizationID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetCurrentTermByOrganizationIDCacheKey(organizationID))
}
