package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedHolidayService interface {
	GetStudentCheckInCacheKey(ctx context.Context, studentID string) (map[string]interface{}, error)
	GetUserCheckInCacheKey(ctx context.Context, userID string) (map[string]interface{}, error)
	GetStudentCheckOutCacheKey(ctx context.Context, studentID string) (map[string]interface{}, error)
	GetUserCheckOutCacheKey(ctx context.Context, userID string) (map[string]interface{}, error)
}

type cachedHolidayService struct {
	cache *cache.RedisCache
}

func NewCachedHolidayService(cache *cache.RedisCache) CachedHolidayService {
	return &cachedHolidayService{
		cache: cache,
	}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedHolidayService) GetStudentCheckInCacheKey(ctx context.Context, studentID string) (map[string]interface{}, error) {
	if studentID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.StudentCheckInCacheKey(studentID))
}

func (c *cachedHolidayService) GetUserCheckInCacheKey(ctx context.Context, userID string) (map[string]interface{}, error) {
	if userID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.UserCheckInCacheKey(userID))
}

func (c *cachedHolidayService) GetStudentCheckOutCacheKey(ctx context.Context, studentID string) (map[string]interface{}, error) {
	if studentID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.StudentCheckOutCacheKey(studentID))
}

func (c *cachedHolidayService) GetUserCheckOutCacheKey(ctx context.Context, userID string) (map[string]interface{}, error) {
	if userID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.UserCheckOutCacheKey(userID))
}
