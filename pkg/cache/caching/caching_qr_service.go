package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingQRService interface {
	SetQRCode(ctx context.Context, qrCodeID string, data interface{}) error
	InvalidateQRCode(ctx context.Context, qrCodeID string) error
}

type cachingQRService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingQRService(cache *cache.RedisCache, defaultTTL int) CachingQRService {
	return &cachingQRService{
		cache:      cache,
		defaultTTL: defaultTTL,
	}
}

func (s *cachingQRService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingQRService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingQRService) SetQRCode(ctx context.Context, qrCodeID string, data interface{}) error {
	if qrCodeID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.GetQRCodeCacheKey(qrCodeID), data)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingQRService) InvalidateQRCode(ctx context.Context, qrCodeID string) error {
	if qrCodeID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetQRCodeCacheKey(qrCodeID))
}
