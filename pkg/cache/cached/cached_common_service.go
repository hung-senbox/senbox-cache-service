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
// === GET CACHE ===
// ========================

func (c *cachedCommonService) GetLanguageById(ctx context.Context, languageID uint) (map[string]interface{}, error) {
	if languageID == 0 {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetLanguageByIdCacheKey(languageID))
}
