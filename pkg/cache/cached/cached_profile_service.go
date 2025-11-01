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
