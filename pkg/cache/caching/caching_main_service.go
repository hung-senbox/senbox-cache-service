package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	mainservicedto "github.com/hung-senbox/senbox-cache-service/pkg/cache/dto/main_service_dto"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingMainService interface {
	SetUserCache(ctx context.Context, user *mainservicedto.UserResponse) error
	SetStudentCache(ctx context.Context, student *mainservicedto.StudentResponse) error
	SetTeacherCache(ctx context.Context, teacher *mainservicedto.TeacherResponse) error
	SetStaffCache(ctx context.Context, staff *mainservicedto.StaffResponse) error
	SetParentCache(ctx context.Context, parent *mainservicedto.ParentResponse) error
	SetChildCache(ctx context.Context, child *mainservicedto.ChildResponse) error
	SetTeacherByUserAndOrgCacheKey(ctx context.Context, userID, orgID string, teacher *mainservicedto.TeacherResponse) error
	SetStaffByUserAndOrgCacheKey(ctx context.Context, userID, orgID string, staff *mainservicedto.StaffResponse) error
	SetUserByTeacherCacheKey(ctx context.Context, teacherID string, user *mainservicedto.UserResponse) error
	SetParentByUserCacheKey(ctx context.Context, userID string, parent *mainservicedto.ParentResponse) error

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

func (s *cachingMainService) SetUserCache(ctx context.Context, user *mainservicedto.UserResponse) error {
	if user == nil || user.ID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.UserCacheKey(user.ID), user)
}

func (s *cachingMainService) SetStudentCache(ctx context.Context, student *mainservicedto.StudentResponse) error {
	if student == nil || student.StudentID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.StudentCacheKey(student.StudentID), student)
}

func (s *cachingMainService) SetTeacherCache(ctx context.Context, teacher *mainservicedto.TeacherResponse) error {
	if teacher == nil || teacher.TeacherID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.TeacherCacheKey(teacher.TeacherID), teacher)
}

func (s *cachingMainService) SetParentCache(ctx context.Context, parent *mainservicedto.ParentResponse) error {
	if parent == nil || parent.ParentID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.ParentCacheKey(parent.ParentID), parent)
}

func (s *cachingMainService) SetStaffCache(ctx context.Context, staff *mainservicedto.StaffResponse) error {
	if staff == nil || staff.StaffID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.StaffCacheKey(staff.StaffID), staff)
}

func (s *cachingMainService) SetChildCache(ctx context.Context, child *mainservicedto.ChildResponse) error {
	if child == nil || child.ChildID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.ChildCacheKey(child.ChildID), child)
}

func (s *cachingMainService) SetTeacherByUserAndOrgCacheKey(ctx context.Context, userID, orgID string, teacher *mainservicedto.TeacherResponse) error {
	if teacher == nil || userID == "" || orgID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.TeacherByUserAndOrgCacheKey(userID, orgID), teacher)
}

func (s *cachingMainService) SetStaffByUserAndOrgCacheKey(ctx context.Context, userID, orgID string, staff *mainservicedto.StaffResponse) error {
	if staff == nil || userID == "" || orgID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.StaffByUserAndOrgCacheKey(userID, orgID), staff)
}

func (s *cachingMainService) SetUserByTeacherCacheKey(ctx context.Context, teacherID string, user *mainservicedto.UserResponse) error {
	if user == nil || teacherID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.UserByTeacherCacheKey(teacherID), user)
}

func (s *cachingMainService) SetParentByUserCacheKey(ctx context.Context, userID string, parent *mainservicedto.ParentResponse) error {
	if parent == nil || userID == "" {
		return nil
	}
	return s.setByKey(ctx, keys.ParentByUserCacheKey(userID), parent)
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
