package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedMainGateway interface {
	GetUserCache(ctx context.Context, userID string) (map[string]interface{}, error)
	GetStudentCache(ctx context.Context, studentID string) (map[string]interface{}, error)
	GetTeacherCache(ctx context.Context, teacherID string) (map[string]interface{}, error)
	GetStaffCache(ctx context.Context, staffID string) (map[string]interface{}, error)
	GetParentCache(ctx context.Context, parentID string) (map[string]interface{}, error)
	GetChildCache(ctx context.Context, childID string) (map[string]interface{}, error)
	GetTeacherByUserAndOrgCache(ctx context.Context, userID, orgID string) (map[string]interface{}, error)
	GetStaffByUserAndOrgCache(ctx context.Context, userID, orgID string) (map[string]interface{}, error)
	GetUserByTeacherCache(ctx context.Context, teacherID string) (map[string]interface{}, error)
	GetParentByUserCache(ctx context.Context, userID string) (map[string]interface{}, error)

	// ======================== Image Cache ========================
	GetUserImageCache(ctx context.Context, userID string) (map[string]interface{}, error)
	GetStudentImageCache(ctx context.Context, studentID string) (map[string]interface{}, error)
	GetTeacherImageCache(ctx context.Context, teacherID string) (map[string]interface{}, error)
	GetStaffImageCache(ctx context.Context, staffID string) (map[string]interface{}, error)
	GetParentImageCache(ctx context.Context, parentID string) (map[string]interface{}, error)
	GetChildImageCache(ctx context.Context, childID string) (map[string]interface{}, error)

	// ======================== Image by key Cache ========================
	GetImageByKeyCache(ctx context.Context, key string) (map[string]interface{}, error)
}

type cachedMainService struct {
	cache *cache.RedisCache
}

func NewCachedMainGateway(cache *cache.RedisCache) CachedMainGateway {
	return &cachedMainService{
		cache: cache,
	}
}

// ========================
// === HELPER METHODS ===
// ========================

func (c *cachedMainService) getCache(
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
// === GET CACHE ===
// ========================

func (c *cachedMainService) GetUserCache(ctx context.Context, userID string) (map[string]interface{}, error) {
	if userID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.UserCacheKey(userID))
}

func (c *cachedMainService) GetStudentCache(ctx context.Context, studentID string) (map[string]interface{}, error) {
	if studentID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.StudentCacheKey(studentID))
}

func (c *cachedMainService) GetTeacherCache(ctx context.Context, teacherID string) (map[string]interface{}, error) {
	if teacherID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.TeacherCacheKey(teacherID))
}

func (c *cachedMainService) GetStaffCache(ctx context.Context, staffID string) (map[string]interface{}, error) {
	if staffID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.StaffCacheKey(staffID))
}

func (c *cachedMainService) GetParentCache(ctx context.Context, parentID string) (map[string]interface{}, error) {
	if parentID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.ParentCacheKey(parentID))
}

func (c *cachedMainService) GetChildCache(ctx context.Context, childID string) (map[string]interface{}, error) {
	if childID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.ChildCacheKey(childID))
}

func (c *cachedMainService) GetTeacherByUserAndOrgCache(ctx context.Context, userID, orgID string) (map[string]interface{}, error) {
	if userID == "" || orgID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.TeacherByUserAndOrgCacheKey(userID, orgID))
}

func (c *cachedMainService) GetStaffByUserAndOrgCache(ctx context.Context, userID, orgID string) (map[string]interface{}, error) {
	if userID == "" || orgID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.StaffByUserAndOrgCacheKey(userID, orgID))
}

func (c *cachedMainService) GetUserByTeacherCache(ctx context.Context, teacherID string) (map[string]interface{}, error) {
	if teacherID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.UserByTeacherCacheKey(teacherID))
}

func (c *cachedMainService) GetParentByUserCache(ctx context.Context, userID string) (map[string]interface{}, error) {
	if userID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.ParentByUserCacheKey(userID))
}

// ======================== Image Cache ========================
func (c *cachedMainService) GetUserImageCache(ctx context.Context, userID string) (map[string]interface{}, error) {
	if userID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.UserImageCacheKey(userID))
}

func (c *cachedMainService) GetStudentImageCache(ctx context.Context, studentID string) (map[string]interface{}, error) {
	if studentID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.StudentImageCacheKey(studentID))
}

func (c *cachedMainService) GetTeacherImageCache(ctx context.Context, teacherID string) (map[string]interface{}, error) {
	if teacherID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.TeacherImageCacheKey(teacherID))
}

func (c *cachedMainService) GetStaffImageCache(ctx context.Context, staffID string) (map[string]interface{}, error) {
	if staffID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.StaffImageCacheKey(staffID))
}

func (c *cachedMainService) GetParentImageCache(ctx context.Context, parentID string) (map[string]interface{}, error) {
	if parentID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.ParentImageCacheKey(parentID))
}

func (c *cachedMainService) GetChildImageCache(ctx context.Context, childID string) (map[string]interface{}, error) {
	if childID == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.ChildImageCacheKey(childID))
}

// ======================== Image by key Cache ========================
func (c *cachedMainService) GetImageByKeyCache(ctx context.Context, key string) (map[string]interface{}, error) {
	if key == "" {
		return nil, nil
	}
	return c.getCache(ctx, keys.ImageByKeyCacheKey(key))
}
