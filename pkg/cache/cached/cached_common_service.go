package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedCommonGateway interface {
	GetLanguageById(ctx context.Context, languageID uint) (map[string]interface{}, error)
}

type cachedCommonService struct {
	cache *cache.RedisCache
}

func NewCachedCommonGateway(cache *cache.RedisCache) CachedCommonGateway {
	return &cachedCommonService{
		cache: cache,
	}
}

// ========================
// === HELPER METHODS ===
// ========================

func (c *cachedCommonService) getCache(
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

func (c *cachedCommonService) GetLanguageById(ctx context.Context, languageID uint) (map[string]interface{}, error) {
	if languageID == 0 {
		return nil, nil
	}
	return c.getCache(ctx, keys.GetLanguageByIdCacheKey(languageID))
}
