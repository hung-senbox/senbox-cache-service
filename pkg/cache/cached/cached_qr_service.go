package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/qr"
)

type CachedQRService interface {
	GetQRCodeCache(ctx context.Context, qrCodeID string) (map[string]interface{}, error)
	GetLibModeByID(ctx context.Context, libModeID string) ([]qr.LibMode, error)
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

func (c *cachedQRService) GetLibModeByID(ctx context.Context, libModeID string) ([]qr.LibMode, error) {
	if libModeID == "" {
		return nil, nil
	}

	var result []qr.LibMode
	if err := c.cache.Get(ctx, keys.GetLibModeCacheKey(libModeID), &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result, nil

}
