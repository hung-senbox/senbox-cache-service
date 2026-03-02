package caching

import (
	"context"
	"time"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingHolidayService interface {
	SetStudentCheckInCacheKey(ctx context.Context, studentID string, data interface{}) error
	SetTeacherCheckInCacheKey(ctx context.Context, teacherID string, data interface{}) error
	SetStaffCheckInCacheKey(ctx context.Context, staffID string, data interface{}) error
	SetStudentCheckOutCacheKey(ctx context.Context, studentID string, data interface{}) error
	SetTeacherCheckOutCacheKey(ctx context.Context, teacherID string, data interface{}) error
	SetStaffCheckOutCacheKey(ctx context.Context, staffID string, data interface{}) error

	InvalidateStudentCheckInCacheKey(ctx context.Context, studentID string) error
	InvalidateTeacherCheckInCacheKey(ctx context.Context, teacherID string) error
	InvalidateStaffCheckInCacheKey(ctx context.Context, staffID string) error
	InvalidateStudentCheckOutCacheKey(ctx context.Context, studentID string) error
	InvalidateTeacherCheckOutCacheKey(ctx context.Context, teacherID string) error
	InvalidateStaffCheckOutCacheKey(ctx context.Context, staffID string) error
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

func (s *cachingHolidayService) SetTeacherCheckInCacheKey(ctx context.Context, teacherID string, data interface{}) error {
	if teacherID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.TeacherCheckInCacheKey(teacherID), data)
}

func (s *cachingHolidayService) SetStaffCheckInCacheKey(ctx context.Context, staffID string, data interface{}) error {
	if staffID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.StaffCheckInCacheKey(staffID), data)
}

func (s *cachingHolidayService) SetStudentCheckOutCacheKey(ctx context.Context, studentID string, data interface{}) error {
	if studentID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.StudentCheckOutCacheKey(studentID), data)
}

func (s *cachingHolidayService) SetTeacherCheckOutCacheKey(ctx context.Context, teacherID string, data interface{}) error {
	if teacherID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.TeacherCheckOutCacheKey(teacherID), data)
}

func (s *cachingHolidayService) SetStaffCheckOutCacheKey(ctx context.Context, staffID string, data interface{}) error {
	if staffID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.StaffCheckOutCacheKey(staffID), data)
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

func (s *cachingHolidayService) InvalidateTeacherCheckInCacheKey(ctx context.Context, teacherID string) error {
	if teacherID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.TeacherCheckInCacheKey(teacherID))
}

func (s *cachingHolidayService) InvalidateStaffCheckInCacheKey(ctx context.Context, staffID string) error {
	if staffID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StaffCheckInCacheKey(staffID))
}

func (s *cachingHolidayService) InvalidateStudentCheckOutCacheKey(ctx context.Context, studentID string) error {
	if studentID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StudentCheckOutCacheKey(studentID))
}

func (s *cachingHolidayService) InvalidateTeacherCheckOutCacheKey(ctx context.Context, teacherID string) error {
	if teacherID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.TeacherCheckOutCacheKey(teacherID))
}

func (s *cachingHolidayService) InvalidateStaffCheckOutCacheKey(ctx context.Context, staffID string) error {
	if staffID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StaffCheckOutCacheKey(staffID))
}
