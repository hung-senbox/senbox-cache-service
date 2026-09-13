package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/department"
)

type CachingDepartmentService interface {
	SetHierarchyDepartmentCacheKey(ctx context.Context, organizationID string, hierarchyDepartments []*department.Department) error
	InvalidateHierarchyDepartmentCacheKey(ctx context.Context, organizationID string) error
}

type cachingDepartmentService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingDepartmentService(redisCache *cache.RedisCache, defaultTTL int) CachingDepartmentService {
	return &cachingDepartmentService{
		cache:      redisCache,
		defaultTTL: defaultTTL,
	}
}

func (s *cachingDepartmentService) setByKey(ctx context.Context, key string, data interface{}) error {
	if data == nil {
		return nil
	}
	return s.cache.Set(ctx, key, data, s.defaultTTL)
}

func (s *cachingDepartmentService) deleteByKey(ctx context.Context, key string) error {
	return s.cache.Delete(ctx, key)
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingDepartmentService) SetHierarchyDepartmentCacheKey(ctx context.Context, organizationID string, hierarchyDepartments []*department.Department) error {
	if organizationID == "" || hierarchyDepartments == nil {
		return nil
	}
	return s.setByKey(ctx, keys.GetOrganizationHierarchyDepartmentCacheKey(organizationID), hierarchyDepartments)
}

// ========================
// === INVALIDATE CACHE ===
// ========================

func (s *cachingDepartmentService) InvalidateHierarchyDepartmentCacheKey(ctx context.Context, organizationID string) error {
	if organizationID == "" {
		return nil
	}
	return s.deleteByKey(ctx, keys.GetOrganizationHierarchyDepartmentCacheKey(organizationID))
}
