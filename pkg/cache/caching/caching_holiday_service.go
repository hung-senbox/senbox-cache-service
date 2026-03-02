package caching

import (
	"context"
	"time"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingHolidayService interface {
	SetUserCheckInCacheKey(ctx context.Context, userID string, data interface{}) error
	SetStudentCheckInCacheKey(ctx context.Context, studentID string, data interface{}) error
	SetStudentCheckOutCacheKey(ctx context.Context, studentID string, data interface{}) error
	SetUserCheckOutCacheKey(ctx context.Context, userID string, data interface{}) error

	InvalidateStudentCheckInCacheKey(ctx context.Context, studentID string) error
	InvalidateUserCheckOutCacheKey(ctx context.Context, userID string) error
	InvalidateStudentCheckOutCacheKey(ctx context.Context, studentID string) error
}

type cachingHolidayService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingHolidayService(redisCache *cache.RedisCache, defaultTTL int) CachingHolidayService {
	return &cachingHolidayService{
		cache:      redisCache,
		defaultTTL: defaultTTL,
	}
}

func (s *cachingHolidayService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}

	ttl := s.ttlUntilEndOfDaySeconds()
	if ttl <= 0 {
		ttl = s.defaultTTL
	}

	return s.cache.Set(ctx, key, data, ttl)
}

func (s *cachingHolidayService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ttlUntilEndOfDaySeconds returns the remaining seconds to midnight in local time
func (s *cachingHolidayService) ttlUntilEndOfDaySeconds() int {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	return int(endOfDay.Sub(now).Seconds())
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingHolidayService) SetStudentCheckInCacheKey(ctx context.Context, studentID string, data interface{}) error {
	if studentID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.StudentCheckInCacheKey(studentID), data)
}

func (s *cachingHolidayService) SetUserCheckInCacheKey(ctx context.Context, userID string, data interface{}) error {
	if userID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.UserCheckInCacheKey(userID), data)
}

func (s *cachingHolidayService) SetStudentCheckOutCacheKey(ctx context.Context, studentID string, data interface{}) error {
	if studentID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.StudentCheckOutCacheKey(studentID), data)
}

func (s *cachingHolidayService) SetUserCheckOutCacheKey(ctx context.Context, userID string, data interface{}) error {
	if userID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.UserCheckOutCacheKey(userID), data)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingHolidayService) InvalidateStudentCheckInCacheKey(ctx context.Context, studentID string) error {
	if studentID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StudentCheckInCacheKey(studentID))
}

func (s *cachingHolidayService) InvalidateUserCheckInCacheKey(ctx context.Context, userID string) error {
	if userID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.UserCheckInCacheKey(userID))
}

func (s *cachingHolidayService) InvalidateStudentCheckOutCacheKey(ctx context.Context, studentID string) error {
	if studentID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StudentCheckOutCacheKey(studentID))
}

func (s *cachingHolidayService) InvalidateUserCheckOutCacheKey(ctx context.Context, userID string) error {
	if userID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.UserCheckOutCacheKey(userID))
}
