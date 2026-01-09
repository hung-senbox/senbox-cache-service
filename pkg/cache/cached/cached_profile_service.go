package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedProfileGateway interface {
	GetUserCode(ctx context.Context, userID string) (string, error)
	GetStudentCode(ctx context.Context, studentID string) (string, error)
	GetTeacherCode(ctx context.Context, teacherID string) (string, error)
	GetStaffCode(ctx context.Context, staffID string) (string, error)
	GetParentCode(ctx context.Context, parentID string) (string, error)
	GetChildCode(ctx context.Context, childID string) (string, error)
	GetDeviceCode(ctx context.Context, deviceID string) (string, error)
	GetOrganizationCode(ctx context.Context, organizationID string) (string, error)
	GetChildEnrollment(ctx context.Context, childID string) (map[string]interface{}, error)

	// language setting cache key
	GetStudentLanguageSetting(ctx context.Context, studentID string) (map[string]interface{}, error)
	GetTeacherLanguageSetting(ctx context.Context, teacherID string) (map[string]interface{}, error)
	GetStaffLanguageSetting(ctx context.Context, staffID string) (map[string]interface{}, error)
	GetParentLanguageSetting(ctx context.Context, parentID string) (map[string]interface{}, error)
	GetChildLanguageSetting(ctx context.Context, childID string) (map[string]interface{}, error)
}

type cachedProfileService struct {
	cache *cache.RedisCache
}

func NewCachedProfileGateway(cache *cache.RedisCache) CachedProfileGateway {
	return &cachedProfileService{
		cache: cache,
	}
}

// ========================
// === HELPER METHODS ===
// ========================

func (c *cachedProfileService) getCode(
	ctx context.Context,
	cacheKey string,
) (string, error) {
	if cacheKey == "" {
		return "", nil
	}

	var cached string
	if err := c.cache.Get(ctx, cacheKey, &cached); err == nil && cached != "" {
		return cached, nil
	}

	return "", nil
}

func (c *cachedProfileService) getCacheJson(
	ctx context.Context,
	cacheKey string,
) (map[string]interface{}, error) {
	if cacheKey == "" {
		return nil, nil
	}

	var result map[string]interface{}
	if err := c.cache.Get(ctx, cacheKey, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

// ========================
// === GET CODE CACHE ===
// ========================

func (c *cachedProfileService) GetUserCode(ctx context.Context, userID string) (string, error) {
	if userID == "" {
		return "", nil
	}
	return c.getCode(ctx, keys.UserCodeCacheKey(userID))
}

func (c *cachedProfileService) GetStudentCode(ctx context.Context, studentID string) (string, error) {
	if studentID == "" {
		return "", nil
	}
	return c.getCode(ctx, keys.StudentCodeCacheKey(studentID))
}

func (c *cachedProfileService) GetTeacherCode(ctx context.Context, teacherID string) (string, error) {
	if teacherID == "" {
		return "", nil
	}
	return c.getCode(ctx, keys.TeacherCodeCacheKey(teacherID))
}

func (c *cachedProfileService) GetStaffCode(ctx context.Context, staffID string) (string, error) {
	if staffID == "" {
		return "", nil
	}
	return c.getCode(ctx, keys.StaffCodeCacheKey(staffID))
}

func (c *cachedProfileService) GetParentCode(ctx context.Context, parentID string) (string, error) {
	if parentID == "" {
		return "", nil
	}
	return c.getCode(ctx, keys.ParentCodeCacheKey(parentID))
}

func (c *cachedProfileService) GetChildCode(ctx context.Context, childID string) (string, error) {
	if childID == "" {
		return "", nil
	}
	return c.getCode(ctx, keys.ChildCodeCacheKey(childID))
}

func (c *cachedProfileService) GetDeviceCode(ctx context.Context, deviceID string) (string, error) {
	if deviceID == "" {
		return "", nil
	}
	return c.getCode(ctx, keys.DeviceCodeCacheKey(deviceID))
}

func (c *cachedProfileService) GetOrganizationCode(ctx context.Context, organizationID string) (string, error) {
	if organizationID == "" {
		return "", nil
	}
	return c.getCode(ctx, keys.OrganizationCodeCacheKey(organizationID))
}

func (c *cachedProfileService) GetChildEnrollment(ctx context.Context, childID string) (map[string]interface{}, error) {
	if childID == "" {
		return nil, nil
	}
	return c.getCacheJson(ctx, keys.ChildEnrollmentCodeCacheKey(childID))
}

// ========================
// === GET LANGUAGE SETTING CACHE ===
// ========================

func (c *cachedProfileService) GetStudentLanguageSetting(ctx context.Context, studentID string) (map[string]interface{}, error) {
	if studentID == "" {
		return nil, nil
	}
	return c.getCacheJson(ctx, keys.StudentLanguageSettingCacheKey(studentID))
}

func (c *cachedProfileService) GetTeacherLanguageSetting(ctx context.Context, teacherID string) (map[string]interface{}, error) {
	if teacherID == "" {
		return nil, nil
	}
	return c.getCacheJson(ctx, keys.TeacherLanguageSettingCacheKey(teacherID))
}

func (c *cachedProfileService) GetStaffLanguageSetting(ctx context.Context, staffID string) (map[string]interface{}, error) {
	if staffID == "" {
		return nil, nil
	}
	return c.getCacheJson(ctx, keys.StaffLanguageSettingCacheKey(staffID))
}

func (c *cachedProfileService) GetParentLanguageSetting(ctx context.Context, parentID string) (map[string]interface{}, error) {
	if parentID == "" {
		return nil, nil
	}
	return c.getCacheJson(ctx, keys.ParentLanguageSettingCacheKey(parentID))
}

func (c *cachedProfileService) GetChildLanguageSetting(ctx context.Context, childID string) (map[string]interface{}, error) {
	if childID == "" {
		return nil, nil
	}
	return c.getCacheJson(ctx, keys.ChildLanguageSettingCacheKey(childID))
}
