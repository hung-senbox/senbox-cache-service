package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedReportService interface {
	GetStudentReportPercentage(ctx context.Context, studentID string) (float32, error)
	GetTeacherReportPercentage(ctx context.Context, teacherID string) (float32, error)
}

type cachedReportService struct {
	cache *cache.RedisCache
}

func NewCachedReportService(cache *cache.RedisCache) CachedReportService {
	return &cachedReportService{
		cache: cache,
	}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedReportService) GetStudentReportPercentage(ctx context.Context, studentID string) (float32, error) {
	if studentID == "" {
		return 0, nil
	}
	return getCacheFloat(c.cache, ctx, keys.StudentReportPercentage(studentID))
}

func (c *cachedReportService) GetTeacherReportPercentage(ctx context.Context, teacherID string) (float32, error) {
	if teacherID == "" {
		return 0, nil
	}
	return getCacheFloat(c.cache, ctx, keys.TeacherReportPercentage(teacherID))
}
