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
	cache cache.Cache
	ttl   int
}

func NewCachedProfileGateway(cache cache.Cache, ttl int) CachedProfileGateway {
	return &cachedProfileService{
		cache: cache,
		ttl:   ttl,
	}
}

// -------------------------
// Generic helper
// -------------------------

func (c *cachedProfileService) getCode(
	ctx context.Context,
	cacheKey string,
) (string, error) {
	var cached string
	if err := c.cache.Get(ctx, cacheKey, &cached); err == nil && cached != "" {
		return cached, nil
	}

	return "", nil
}

// -------------------------
// Public methods
// -------------------------

func (c *cachedProfileService) GetStudentCode(ctx context.Context, studentID string) (string, error) {
	return c.getCode(ctx, keys.StudentCodeCacheKey(studentID))
}

func (c *cachedProfileService) GetTeacherCode(ctx context.Context, teacherID string) (string, error) {
	return c.getCode(ctx, keys.TeacherCodeCacheKey(teacherID))
}

func (c *cachedProfileService) GetParentCode(ctx context.Context, parentID string) (string, error) {
	return c.getCode(ctx, keys.ParentCodeCacheKey(parentID))
}

func (c *cachedProfileService) GetStaffCode(ctx context.Context, staffID string) (string, error) {
	return c.getCode(ctx, keys.StaffCodeCacheKey(staffID))
}

func (c *cachedProfileService) GetChildCode(ctx context.Context, childID string) (string, error) {
	return c.getCode(ctx, keys.ChildCodeCacheKey(childID))
}

func (c *cachedProfileService) GetUserCode(ctx context.Context, userID string) (string, error) {
	return c.getCode(ctx, keys.UserCodeCacheKey(userID))
}
