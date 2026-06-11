package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/media"
)

type CachedCountService interface {
	GetMediaPortalCountByStudentID(ctx context.Context, studentID string) ([]media.CountMediaStudentPortal, error)
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
