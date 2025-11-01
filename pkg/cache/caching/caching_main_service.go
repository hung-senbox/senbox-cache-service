package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingMainService interface {
	SetUserCache(ctx context.Context, userID string, data interface{}) error
	SetStudentCache(ctx context.Context, studentID string, data interface{}) error
	SetTeacherCache(ctx context.Context, teacherID string, data interface{}) error
	SetStaffCache(ctx context.Context, staffID string, data interface{}) error
	SetParentCache(ctx context.Context, parentID string, data interface{}) error
	SetChildCache(ctx context.Context, childID string, data interface{}) error
	SetTeacherByUserAndOrgCacheKey(ctx context.Context, userID, orgID string, data interface{}) error
	SetStaffByUserAndOrgCacheKey(ctx context.Context, userID, orgID string, data interface{}) error
	SetUserByTeacherCacheKey(ctx context.Context, teacherID string, data interface{}) error
	SetParentByUserCacheKey(ctx context.Context, userID string, data interface{}) error

	InvalidateUserCache(ctx context.Context, userID string) error
	InvalidateTeacherCache(ctx context.Context, teacherID string) error
	InvalidateParentCache(ctx context.Context, parentID string) error
	InvalidateStaffCache(ctx context.Context, staffID string) error
	InvalidateChildCache(ctx context.Context, childID string) error
	InvalidateStudentCache(ctx context.Context, studentID string) error
	InvalidateUserByTeacherCacheKey(ctx context.Context, teacherID string) error
	InvalidateParentByUserCacheKey(ctx context.Context, userID string) error
	InvalidateTeacherByUserAndOrgCacheKey(ctx context.Context, userID, orgID string) error
	InvalidateStaffByUserAndOrgCacheKey(ctx context.Context, userID, orgID string) error
}

type cachingMainService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingMainService(redisCache *cache.RedisCache, ttlSeconds int) CachingMainService {
	return &cachingMainService{
		cache:      redisCache,
		defaultTTL: ttlSeconds,
	}
}

func (s *cachingMainService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingMainService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingMainService) SetUserCache(ctx context.Context, userID string, data interface{}) error {
	if userID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.UserCacheKey(userID), data)
}

func (s *cachingMainService) SetStudentCache(ctx context.Context, studentID string, data interface{}) error {
	if studentID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.StudentCacheKey(studentID), data)
}

func (s *cachingMainService) SetTeacherCache(ctx context.Context, teacherID string, data interface{}) error {
	if teacherID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.TeacherCacheKey(teacherID), data)
}

func (s *cachingMainService) SetParentCache(ctx context.Context, parentID string, data interface{}) error {
	if parentID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.ParentCacheKey(parentID), data)
}

func (s *cachingMainService) SetStaffCache(ctx context.Context, staffID string, data interface{}) error {
	if staffID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.StaffCacheKey(staffID), data)
}

func (s *cachingMainService) SetChildCache(ctx context.Context, childID string, data interface{}) error {
	if childID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.ChildCacheKey(childID), data)
}

func (s *cachingMainService) SetTeacherByUserAndOrgCacheKey(ctx context.Context, userID, orgID string, data interface{}) error {
	if userID == "" || orgID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.TeacherByUserAndOrgCacheKey(userID, orgID), data)
}

func (s *cachingMainService) SetStaffByUserAndOrgCacheKey(ctx context.Context, userID, orgID string, data interface{}) error {
	if userID == "" || orgID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.StaffByUserAndOrgCacheKey(userID, orgID), data)
}

func (s *cachingMainService) SetUserByTeacherCacheKey(ctx context.Context, teacherID string, data interface{}) error {
	if teacherID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.UserByTeacherCacheKey(teacherID), data)
}

func (s *cachingMainService) SetParentByUserCacheKey(ctx context.Context, userID string, data interface{}) error {
	if userID == "" || data == nil {
		return nil
	}
	return s.setByKey(ctx, keys.ParentByUserCacheKey(userID), data)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingMainService) InvalidateUserCache(ctx context.Context, userID string) error {
	if userID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.UserCacheKey(userID))
}

func (s *cachingMainService) InvalidateStudentCache(ctx context.Context, studentID string) error {
	if studentID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StudentCacheKey(studentID))
}

func (s *cachingMainService) InvalidateTeacherCache(ctx context.Context, teacherID string) error {
	if teacherID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.TeacherCacheKey(teacherID))
}

func (s *cachingMainService) InvalidateParentCache(ctx context.Context, parentID string) error {
	if parentID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.ParentCacheKey(parentID))
}

func (s *cachingMainService) InvalidateStaffCache(ctx context.Context, staffID string) error {
	if staffID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StaffCacheKey(staffID))
}

func (s *cachingMainService) InvalidateChildCache(ctx context.Context, childID string) error {
	if childID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.ChildCacheKey(childID))
}

func (s *cachingMainService) InvalidateUserByTeacherCacheKey(ctx context.Context, teacherID string) error {
	if teacherID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.UserByTeacherCacheKey(teacherID))
}

func (s *cachingMainService) InvalidateParentByUserCacheKey(ctx context.Context, userID string) error {
	if userID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.ParentByUserCacheKey(userID))
}

func (s *cachingMainService) InvalidateTeacherByUserAndOrgCacheKey(ctx context.Context, userID, orgID string) error {
	if userID == "" || orgID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.TeacherByUserAndOrgCacheKey(userID, orgID))
}

func (s *cachingMainService) InvalidateStaffByUserAndOrgCacheKey(ctx context.Context, userID, orgID string) error {
	if userID == "" || orgID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StaffByUserAndOrgCacheKey(userID, orgID))
}
