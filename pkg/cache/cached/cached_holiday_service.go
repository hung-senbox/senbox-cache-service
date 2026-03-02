package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedHolidayService interface {
	GetStudentCheckInCacheKey(ctx context.Context, studentID string) (map[string]interface{}, error)
	GetTeacherCheckInCacheKey(ctx context.Context, teacherID string) (map[string]interface{}, error)
	GetStaffCheckInCacheKey(ctx context.Context, staffID string) (map[string]interface{}, error)
	GetStudentCheckOutCacheKey(ctx context.Context, studentID string) (map[string]interface{}, error)
	GetTeacherCheckOutCacheKey(ctx context.Context, teacherID string) (map[string]interface{}, error)
	GetStaffCheckOutCacheKey(ctx context.Context, staffID string) (map[string]interface{}, error)
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

func (c *cachedHolidayService) GetTeacherCheckInCacheKey(ctx context.Context, teacherID string) (map[string]interface{}, error) {
	if teacherID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.TeacherCheckInCacheKey(teacherID))
}

func (c *cachedHolidayService) GetStaffCheckInCacheKey(ctx context.Context, staffID string) (map[string]interface{}, error) {
	if staffID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.StaffCheckInCacheKey(staffID))
}

func (c *cachedHolidayService) GetStudentCheckOutCacheKey(ctx context.Context, studentID string) (map[string]interface{}, error) {
	if studentID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.StudentCheckOutCacheKey(studentID))
}

func (c *cachedHolidayService) GetTeacherCheckOutCacheKey(ctx context.Context, teacherID string) (map[string]interface{}, error) {
	if teacherID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.TeacherCheckOutCacheKey(teacherID))
}

func (c *cachedHolidayService) GetStaffCheckOutCacheKey(ctx context.Context, staffID string) (map[string]interface{}, error) {
	if staffID == "" {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.StaffCheckOutCacheKey(staffID))
}
