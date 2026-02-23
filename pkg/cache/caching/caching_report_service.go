package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingReportService interface {
	// Report percentage
	SetStudentReportPercentage(ctx context.Context, studentID string, percentage float32) error
	SetTeacherReportPercentage(ctx context.Context, teacherID string, percentage float32) error
	// Report percentage
	InvalidateStudentReportPercentage(ctx context.Context, studentID string) error
	InvalidateTeacherReportPercentage(ctx context.Context, teacherID string) error
}

type cachingReportService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingReportService(redisCache *cache.RedisCache, defaultTTL int) CachingReportService {
	return &cachingReportService{
		cache:      redisCache,
		defaultTTL: defaultTTL,
	}
}

func (s *cachingReportService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingReportService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingReportService) SetStudentReportPercentage(ctx context.Context, studentID string, percentage float32) error {
	if studentID == "" || percentage == 0 {
		return nil
	}
	return s.setByKey(ctx, keys.StudentReportPercentage(studentID), percentage)
}

func (s *cachingReportService) SetTeacherReportPercentage(ctx context.Context, teacherID string, percentage float32) error {
	if teacherID == "" || percentage == 0 {
		return nil
	}
	return s.setByKey(ctx, keys.TeacherReportPercentage(teacherID), percentage)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingReportService) InvalidateStudentReportPercentage(ctx context.Context, studentID string) error {
	if studentID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StudentReportPercentage(studentID))
}

func (s *cachingReportService) InvalidateTeacherReportPercentage(ctx context.Context, teacherID string) error {
	if teacherID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.TeacherReportPercentage(teacherID))
}
