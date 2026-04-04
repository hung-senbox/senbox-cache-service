package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedCommonGateway interface {
	// ======================== Language Cache ========================
	GetLanguageById(ctx context.Context, languageID uint) (map[string]interface{}, error)

	// ======================== Occupation Cache ========================
	GetOccupationByIdAndLanguageId(ctx context.Context, occupationID string, languageID uint) (map[string]interface{}, error)
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

func (c *cachedCommonService) GetOccupationByIdAndLanguageId(ctx context.Context, occupationID string, languageID uint) (map[string]interface{}, error) {
	if occupationID == "" || languageID == 0 {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetOccupationByIdAndLanguageIdCacheKey(occupationID, languageID))
}
