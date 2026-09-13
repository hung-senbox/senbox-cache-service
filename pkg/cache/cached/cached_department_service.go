package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/department"
)

type CachedDepartmentService interface {
	GetHierarchyDepartmentCacheKey(ctx context.Context, organizationID string) ([]*department.Department, error)
	GetAllLeaderAndStaffsByOrganizationID(ctx context.Context, organizationID string) ([]*department.Leader, []*department.Staff, error)
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

func (c *cachedDepartmentService) GetAllLeaderAndStaffsByOrganizationID(ctx context.Context, organizationID string) ([]*department.Leader, []*department.Staff, error) {
	if organizationID == "" {
		return nil, nil, nil
	}

	departments, err := c.GetHierarchyDepartmentCacheKey(ctx, organizationID)
	if err != nil {
		return nil, nil, err
	}
	if len(departments) == 0 {
		return nil, nil, nil
	}

	leaders := make([]*department.Leader, 0)
	staffs := make([]*department.Staff, 0)
	seenLeaders := make(map[string]struct{})
	seenStaffs := make(map[string]struct{})

	for _, dept := range departments {
		if dept == nil {
			continue
		}

		if dept.Leader.OwnerID != "" {
			if _, ok := seenLeaders[dept.Leader.OwnerID]; !ok {
				seenLeaders[dept.Leader.OwnerID] = struct{}{}
				leader := dept.Leader
				leaders = append(leaders, &leader)
			}
		}

		for i := range dept.Staffs {
			staff := dept.Staffs[i]
			if staff.OwnerID == "" {
				continue
			}
			if _, ok := seenStaffs[staff.OwnerID]; ok {
				continue
			}
			seenStaffs[staff.OwnerID] = struct{}{}
			staffs = append(staffs, &staff)
		}
	}

	if len(leaders) == 0 && len(staffs) == 0 {
		return nil, nil, nil
	}

	return leaders, staffs, nil
}
