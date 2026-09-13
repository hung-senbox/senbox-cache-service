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
	CheckOwnerIsInHierarchyDepartment(ctx context.Context, organizationID string, ownerId []string) (bool, error)
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

func (c *cachedDepartmentService) CheckOwnerIsInHierarchyDepartment(ctx context.Context, organizationID string, ownerId []string) (bool, error) {
	if organizationID == "" || len(ownerId) == 0 {
		return false, nil
	}

	leaders, staffs, err := c.GetAllLeaderAndStaffsByOrganizationID(ctx, organizationID)
	if err != nil {
		return false, err
	}

	hierarchyOwners := make(map[string]struct{}, len(leaders)+len(staffs))
	for _, leader := range leaders {
		if leader != nil && leader.OwnerID != "" {
			hierarchyOwners[leader.OwnerID] = struct{}{}
		}
	}
	for _, staff := range staffs {
		if staff != nil && staff.OwnerID != "" {
			hierarchyOwners[staff.OwnerID] = struct{}{}
		}
	}

	hasOwnerID := false
	for _, id := range ownerId {
		if id == "" {
			continue
		}
		hasOwnerID = true
		if _, ok := hierarchyOwners[id]; !ok {
			return false, nil
		}
	}

	return hasOwnerID, nil
}
