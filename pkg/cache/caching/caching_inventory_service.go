package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingInventoryService interface {
	SetShelfSlotById(ctx context.Context, shelfSlotID string, data interface{}) error
}

type cachingInventoryService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingInventoryService(cache *cache.RedisCache, defaultTTL int) CachingInventoryService {
	return &cachingInventoryService{cache: cache, defaultTTL: defaultTTL}
}

func (s *cachingInventoryService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingInventoryService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingInventoryService) SetShelfSlotById(ctx context.Context, shelfSlotID string, data interface{}) error {
	if shelfSlotID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.GetShelfSlotByIdCacheKey(shelfSlotID), data)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingInventoryService) InvalidateShelfSlotById(ctx context.Context, shelfSlotID string) error {
	if shelfSlotID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetShelfSlotByIdCacheKey(shelfSlotID))
}
