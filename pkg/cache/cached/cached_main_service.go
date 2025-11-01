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
}

type cachedMainService struct {
	cache cache.RedisCache
}

func NewCachedMainGateway(cache cache.RedisCache) CachedMainGateway {
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
