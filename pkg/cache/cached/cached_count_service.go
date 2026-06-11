package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/media"
)

type CachedCountService interface {
	GetMediaPortalCountByStudentID(ctx context.Context, studentID string) ([]media.CountMediaStudentPortal, error)
	GetMediaPortalCountByStudentIds(ctx context.Context, studentIDs []string) (map[string][]media.CountMediaStudentPortal, error)
}

type cachedCountService struct {
	cache *cache.RedisCache
}

func NewCachedCountService(cache *cache.RedisCache) CachedCountService {
	return &cachedCountService{cache: cache}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedCountService) GetMediaPortalCountByStudentID(ctx context.Context, studentID string) ([]media.CountMediaStudentPortal, error) {
	if studentID == "" {
		return nil, nil
	}
	var result []media.CountMediaStudentPortal
	if err := c.cache.Get(ctx, keys.GetCountCacheKey(studentID), &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func (c *cachedCountService) GetMediaPortalCountByStudentIds(ctx context.Context, studentIDs []string) (map[string][]media.CountMediaStudentPortal, error) {
	if len(studentIDs) == 0 {
		return nil, nil
	}

	result := make(map[string][]media.CountMediaStudentPortal, len(studentIDs))
	for _, studentID := range studentIDs {
		if studentID == "" {
			continue
		}

		studentPortalCounts, err := c.GetMediaPortalCountByStudentID(ctx, studentID)
		if err != nil {
			return nil, err
		}
		if len(studentPortalCounts) == 0 {
			continue
		}

		result[studentID] = studentPortalCounts
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
