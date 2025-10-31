package caching

import (
	"context"

	"github.com/hung-senbox/cache-service/pkg/cache"
	keys "github.com/hung-senbox/cache-service/pkg/cache/keys_cache"
)

type CachingProfileService interface {
	SetUserCode(ctx context.Context, userID, code string) error
	SetStudentCode(ctx context.Context, studentID, code string) error
	SetTeacherCode(ctx context.Context, teacherID, code string) error
	SetStaffCode(ctx context.Context, staffID, code string) error
	SetParentCode(ctx context.Context, parentID, code string) error
	SetChildCode(ctx context.Context, childID, code string) error

	InvalidateUserCode(ctx context.Context, userID string) error
	InvalidateStudentCode(ctx context.Context, studentID string) error
	InvalidateTeacherCode(ctx context.Context, teacherID string) error
	InvalidateStaffCode(ctx context.Context, staffID string) error
	InvalidateParentCode(ctx context.Context, parentID string) error
	InvalidateChildCode(ctx context.Context, childID string) error
}

type cachingProfileService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingService(redisCache *cache.RedisCache, ttlSeconds int) CachingProfileService {
	return &cachingProfileService{
		cache:      redisCache,
		defaultTTL: ttlSeconds,
	}
}

func (s *cachingProfileService) setByKey(ctx context.Context, key, code string) error {
	return s.cache.Set(ctx, key, code, s.defaultTTL)
}

func (s *cachingProfileService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CODE CACHE ===
// ========================
func (s *cachingProfileService) SetUserCode(ctx context.Context, userID, code string) error {
	key := keys.UserCodeCacheKey(userID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetStudentCode(ctx context.Context, studentID, code string) error {
	key := keys.StudentCodeCacheKey(studentID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetTeacherCode(ctx context.Context, teacherID, code string) error {
	key := keys.TeacherCodeCacheKey(teacherID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetStaffCode(ctx context.Context, staffID, code string) error {
	key := keys.StaffCodeCacheKey(staffID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetParentCode(ctx context.Context, parentID, code string) error {
	key := keys.ParentCodeCacheKey(parentID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetChildCode(ctx context.Context, childID, code string) error {
	key := keys.ChildCodeCacheKey(childID)
	return s.setByKey(ctx, key, code)
}

// ========================
// === INVALIDATE CACHE ===
// ========================
func (s *cachingProfileService) InvalidateUserCode(ctx context.Context, userID string) error {
	return s.deleteByKey(ctx, keys.UserCodeCacheKey(userID))
}

func (s *cachingProfileService) InvalidateStudentCode(ctx context.Context, studentID string) error {
	return s.deleteByKey(ctx, keys.StudentCodeCacheKey(studentID))
}

func (s *cachingProfileService) InvalidateTeacherCode(ctx context.Context, teacherID string) error {
	return s.deleteByKey(ctx, keys.TeacherCodeCacheKey(teacherID))
}

func (s *cachingProfileService) InvalidateStaffCode(ctx context.Context, staffID string) error {
	return s.deleteByKey(ctx, keys.StaffCodeCacheKey(staffID))
}

func (s *cachingProfileService) InvalidateParentCode(ctx context.Context, parentID string) error {
	return s.deleteByKey(ctx, keys.ParentCodeCacheKey(parentID))
}

func (s *cachingProfileService) InvalidateChildCode(ctx context.Context, childID string) error {
	return s.deleteByKey(ctx, keys.ChildCodeCacheKey(childID))
}
