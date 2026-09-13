package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/department"
)

type CachedDepartmentService interface {
	GetHierarchyDepartmentCacheKey(ctx context.Context, organizationID string) ([]*department.Department, error)
}

type cachedDepartmentService struct {
	cache *cache.RedisCache
}

func NewCachedDepartmentService(cache *cache.RedisCache) CachedDepartmentService {
	return &cachedDepartmentService{cache: cache}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedDepartmentService) GetHierarchyDepartmentCacheKey(ctx context.Context, organizationID string) ([]*department.Department, error) {
	if organizationID == "" {
		return nil, nil
	}

	var result []*department.Department
	if err := c.cache.Get(ctx, keys.GetOrganizationHierarchyDepartmentCacheKey(organizationID), &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
