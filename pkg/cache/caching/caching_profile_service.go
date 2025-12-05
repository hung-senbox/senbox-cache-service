package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachingProfileService interface {
	SetUserCode(ctx context.Context, userID, code string) error
	SetStudentCode(ctx context.Context, studentID, code string) error
	SetTeacherCode(ctx context.Context, teacherID, code string) error
	SetStaffCode(ctx context.Context, staffID, code string) error
	SetParentCode(ctx context.Context, parentID, code string) error
	SetChildCode(ctx context.Context, childID, code string) error
	SetDeviceCode(ctx context.Context, deviceID, code string) error
	SetOrganizationCode(ctx context.Context, organizationID, code string) error
	SetChildEnrollmentCode(ctx context.Context, childID string, enrollmentData map[string]interface{}) error

	InvalidateUserCode(ctx context.Context, userID string) error
	InvalidateStudentCode(ctx context.Context, studentID string) error
	InvalidateTeacherCode(ctx context.Context, teacherID string) error
	InvalidateStaffCode(ctx context.Context, staffID string) error
	InvalidateParentCode(ctx context.Context, parentID string) error
	InvalidateChildCode(ctx context.Context, childID string) error
	InvalidateDeviceCode(ctx context.Context, deviceID string) error
	InvalidateOrganizationCode(ctx context.Context, organizationID string) error
	InvalidateChildEnrollmentCode(ctx context.Context, childID string) error
}

type cachingProfileService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingProfileService(redisCache *cache.RedisCache, ttlSeconds int) CachingProfileService {
	return &cachingProfileService{
		cache:      redisCache,
		defaultTTL: ttlSeconds,
	}
}

func (s *cachingProfileService) setByKey(ctx context.Context, key, code string) error {
	if key == "" || code == "" {
		return nil
	}
	return s.cache.Set(ctx, key, code, s.defaultTTL)
}

func (s *cachingProfileService) setByKeyWithJSON(ctx context.Context, key string, data interface{}) error {
	if key == "" || data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingProfileService) deleteByKey(ctx context.Context, key string) error {
	if key == "" {
		return nil
	}
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CODE CACHE ===
// ========================
func (s *cachingProfileService) SetUserCode(ctx context.Context, userID, code string) error {
	if userID == "" || code == "" {
		return nil
	}
	key := keys.UserCodeCacheKey(userID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetStudentCode(ctx context.Context, studentID, code string) error {
	if studentID == "" || code == "" {
		return nil
	}
	key := keys.StudentCodeCacheKey(studentID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetTeacherCode(ctx context.Context, teacherID, code string) error {
	if teacherID == "" || code == "" {
		return nil
	}
	key := keys.TeacherCodeCacheKey(teacherID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetStaffCode(ctx context.Context, staffID, code string) error {
	if staffID == "" || code == "" {
		return nil
	}
	key := keys.StaffCodeCacheKey(staffID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetParentCode(ctx context.Context, parentID, code string) error {
	if parentID == "" || code == "" {
		return nil
	}
	key := keys.ParentCodeCacheKey(parentID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetChildCode(ctx context.Context, childID, code string) error {
	if childID == "" || code == "" {
		return nil
	}
	key := keys.ChildCodeCacheKey(childID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetDeviceCode(ctx context.Context, deviceID, code string) error {
	if deviceID == "" || code == "" {
		return nil
	}
	key := keys.DeviceCodeCacheKey(deviceID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetOrganizationCode(ctx context.Context, organizationID, code string) error {
	if organizationID == "" || code == "" {
		return nil
	}
	key := keys.OrganizationCodeCacheKey(organizationID)
	return s.setByKey(ctx, key, code)
}

func (s *cachingProfileService) SetChildEnrollmentCode(ctx context.Context, childID string, enrollmentData map[string]interface{}) error {
	if childID == "" || enrollmentData == nil {
		return nil
	}
	key := keys.ChildEnrollmentCodeCacheKey(childID)
	return s.setByKeyWithJSON(ctx, key, enrollmentData)
}

// ========================
// === INVALIDATE CACHE ===
// ========================
func (s *cachingProfileService) InvalidateUserCode(ctx context.Context, userID string) error {
	if userID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.UserCodeCacheKey(userID))
}

func (s *cachingProfileService) InvalidateStudentCode(ctx context.Context, studentID string) error {
	if studentID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StudentCodeCacheKey(studentID))
}

func (s *cachingProfileService) InvalidateTeacherCode(ctx context.Context, teacherID string) error {
	if teacherID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.TeacherCodeCacheKey(teacherID))
}

func (s *cachingProfileService) InvalidateStaffCode(ctx context.Context, staffID string) error {
	if staffID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.StaffCodeCacheKey(staffID))
}

func (s *cachingProfileService) InvalidateParentCode(ctx context.Context, parentID string) error {
	if parentID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.ParentCodeCacheKey(parentID))
}

func (s *cachingProfileService) InvalidateChildCode(ctx context.Context, childID string) error {
	if childID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.ChildCodeCacheKey(childID))
}

func (s *cachingProfileService) InvalidateDeviceCode(ctx context.Context, deviceID string) error {
	if deviceID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.DeviceCodeCacheKey(deviceID))
}

func (s *cachingProfileService) InvalidateOrganizationCode(ctx context.Context, organizationID string) error {
	if organizationID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.OrganizationCodeCacheKey(organizationID))
}

func (s *cachingProfileService) InvalidateChildEnrollmentCode(ctx context.Context, childID string) error {
	if childID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.ChildEnrollmentCodeCacheKey(childID))
}
