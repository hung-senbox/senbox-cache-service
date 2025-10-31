package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	mainservicedto "github.com/hung-senbox/senbox-cache-service/pkg/cache/dto/main_service_dto"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedMainGateway interface {
	GetUserCache(ctx context.Context, userID string) (*mainservicedto.UserResponse, error)
	GetStudentCache(ctx context.Context, studentID string) (*mainservicedto.StudentResponse, error)
	GetTeacherCache(ctx context.Context, teacherID string) (*mainservicedto.TeacherResponse, error)
	GetStaffCache(ctx context.Context, staffID string) (*mainservicedto.StaffResponse, error)
	GetParentCache(ctx context.Context, parentID string) (*mainservicedto.ParentResponse, error)
	GetChildCache(ctx context.Context, childID string) (*mainservicedto.ChildResponse, error)
	GetTeacherByUserAndOrgCache(ctx context.Context, userID, orgID string) (*mainservicedto.TeacherResponse, error)
	GetStaffByUserAndOrgCache(ctx context.Context, userID, orgID string) (*mainservicedto.StaffResponse, error)
	GetUserByTeacherCache(ctx context.Context, teacherID string) (*mainservicedto.UserResponse, error)
	GetParentByUserCache(ctx context.Context, userID string) (*mainservicedto.ParentResponse, error)
}

type cachedMainService struct {
	cache cache.Cache
}

func NewCachedMainGateway(cache cache.Cache) CachedMainGateway {
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
	dest interface{},
) error {
	if cacheKey == "" {
		return nil
	}
	return c.cache.Get(ctx, cacheKey, dest)
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedMainService) GetUserCache(ctx context.Context, userID string) (*mainservicedto.UserResponse, error) {
	if userID == "" {
		return nil, nil
	}
	var user mainservicedto.UserResponse
	if err := c.getCache(ctx, keys.UserCacheKey(userID), &user); err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, nil
	}
	return &user, nil
}

func (c *cachedMainService) GetStudentCache(ctx context.Context, studentID string) (*mainservicedto.StudentResponse, error) {
	if studentID == "" {
		return nil, nil
	}
	var student mainservicedto.StudentResponse
	if err := c.getCache(ctx, keys.StudentCacheKey(studentID), &student); err != nil {
		return nil, err
	}
	if student.StudentID == "" {
		return nil, nil
	}
	return &student, nil
}

func (c *cachedMainService) GetTeacherCache(ctx context.Context, teacherID string) (*mainservicedto.TeacherResponse, error) {
	if teacherID == "" {
		return nil, nil
	}
	var teacher mainservicedto.TeacherResponse
	if err := c.getCache(ctx, keys.TeacherCacheKey(teacherID), &teacher); err != nil {
		return nil, err
	}
	if teacher.TeacherID == "" {
		return nil, nil
	}
	return &teacher, nil
}

func (c *cachedMainService) GetStaffCache(ctx context.Context, staffID string) (*mainservicedto.StaffResponse, error) {
	if staffID == "" {
		return nil, nil
	}
	var staff mainservicedto.StaffResponse
	if err := c.getCache(ctx, keys.StaffCacheKey(staffID), &staff); err != nil {
		return nil, err
	}
	if staff.StaffID == "" {
		return nil, nil
	}
	return &staff, nil
}

func (c *cachedMainService) GetParentCache(ctx context.Context, parentID string) (*mainservicedto.ParentResponse, error) {
	if parentID == "" {
		return nil, nil
	}
	var parent mainservicedto.ParentResponse
	if err := c.getCache(ctx, keys.ParentCacheKey(parentID), &parent); err != nil {
		return nil, err
	}
	if parent.ParentID == "" {
		return nil, nil
	}
	return &parent, nil
}

func (c *cachedMainService) GetChildCache(ctx context.Context, childID string) (*mainservicedto.ChildResponse, error) {
	if childID == "" {
		return nil, nil
	}
	var child mainservicedto.ChildResponse
	if err := c.getCache(ctx, keys.ChildCacheKey(childID), &child); err != nil {
		return nil, err
	}
	if child.ChildID == "" {
		return nil, nil
	}
	return &child, nil
}

func (c *cachedMainService) GetTeacherByUserAndOrgCache(ctx context.Context, userID, orgID string) (*mainservicedto.TeacherResponse, error) {
	if userID == "" || orgID == "" {
		return nil, nil
	}
	var teacher mainservicedto.TeacherResponse
	if err := c.getCache(ctx, keys.TeacherByUserAndOrgCacheKey(userID, orgID), &teacher); err != nil {
		return nil, err
	}
	if teacher.TeacherID == "" {
		return nil, nil
	}
	return &teacher, nil
}

func (c *cachedMainService) GetStaffByUserAndOrgCache(ctx context.Context, userID, orgID string) (*mainservicedto.StaffResponse, error) {
	if userID == "" || orgID == "" {
		return nil, nil
	}
	var staff mainservicedto.StaffResponse
	if err := c.getCache(ctx, keys.StaffByUserAndOrgCacheKey(userID, orgID), &staff); err != nil {
		return nil, err
	}
	if staff.StaffID == "" {
		return nil, nil
	}
	return &staff, nil
}

func (c *cachedMainService) GetUserByTeacherCache(ctx context.Context, teacherID string) (*mainservicedto.UserResponse, error) {
	if teacherID == "" {
		return nil, nil
	}
	var user mainservicedto.UserResponse
	if err := c.getCache(ctx, keys.UserByTeacherCacheKey(teacherID), &user); err != nil {
		return nil, err
	}
	if user.ID == "" {
		return nil, nil
	}
	return &user, nil
}

func (c *cachedMainService) GetParentByUserCache(ctx context.Context, userID string) (*mainservicedto.ParentResponse, error) {
	if userID == "" {
		return nil, nil
	}
	var parent mainservicedto.ParentResponse
	if err := c.getCache(ctx, keys.ParentByUserCacheKey(userID), &parent); err != nil {
		return nil, err
	}
	if parent.ParentID == "" {
		return nil, nil
	}
	return &parent, nil
}
