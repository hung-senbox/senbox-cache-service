package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedQRService interface {
	GetQRCodeCache(ctx context.Context, qrCodeID string) (map[string]interface{}, error)
}

type cachedQRService struct {
	cache *cache.RedisCache
}

func NewCachedQRService(cache *cache.RedisCache) CachedQRService {
	return &cachedQRService{cache: cache}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedQRService) GetQRCodeCache(ctx context.Context, qrCodeID string) (map[string]interface{}, error) {
	if qrCodeID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetQRCodeCacheKey(qrCodeID))
}
