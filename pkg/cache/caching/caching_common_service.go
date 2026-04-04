package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingCommonService interface {
	// ======================== Language Cache ========================
	SetLanguageById(ctx context.Context, languageID uint, data interface{}) error
	InvalidateLanguageById(ctx context.Context, languageID uint) error

	// ======================== Occupation Cache ========================
	SetOccupationByIdAndLanguageId(ctx context.Context, occupationID string, languageID uint, data interface{}) error
	InvalidateOccupationByIdAndLanguageId(ctx context.Context, occupationID string, languageID uint) error
}

type cachingCommonService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingCommonService(redisCache *cache.RedisCache, defaultTTL int) CachingCommonService {
	return &cachingCommonService{
		cache:      redisCache,
		defaultTTL: defaultTTL,
	}
}

func (s *cachingCommonService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingCommonService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingCommonService) SetLanguageById(ctx context.Context, languageID uint, data interface{}) error {
	if languageID == 0 || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.GetLanguageByIdCacheKey(languageID), data)
}

func (s *cachingCommonService) SetOccupationByIdAndLanguageId(ctx context.Context, occupationID string, languageID uint, data interface{}) error {
	if occupationID == "" || languageID == 0 || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.GetOccupationByIdAndLanguageIdCacheKey(occupationID, languageID), data)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingCommonService) InvalidateLanguageById(ctx context.Context, languageID uint) error {
	if languageID == 0 {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetLanguageByIdCacheKey(languageID))
}

func (s *cachingCommonService) InvalidateOccupationByIdAndLanguageId(ctx context.Context, occupationID string, languageID uint) error {
	if occupationID == "" || languageID == 0 {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetOccupationByIdAndLanguageIdCacheKey(occupationID, languageID))
}
