package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/cache/cached"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/helper"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/media"
)

type CachingCountService interface {
	SetMediaPortalCountByStudentID(ctx context.Context, studentID string, studentPortalCounts []media.CountMediaStudentPortal) error // array interface
	SetDefaultCurrentDayMediaPortalCountByStudentIds(ctx context.Context, studentIDs []string) error
	InvalidateMediaPortalCountByStudentID(ctx context.Context, studentID string) error
}

type cachingCountService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingCountService(cache *cache.RedisCache, defaultTTL int) CachingCountService {
	return &cachingCountService{cache: cache, defaultTTL: defaultTTL}
}

// ========================
// === SET CACHE ===
// ========================

func (s *cachingCountService) SetMediaPortalCountByStudentID(ctx context.Context, studentID string, studentPortalCounts []media.CountMediaStudentPortal) error {
	if studentID == "" || studentPortalCounts == nil {
		return nil
	}
	return cached.SetCache(s.cache, ctx, keys.GetCountCacheKey(studentID), studentPortalCounts)
}

func (s *cachingCountService) SetDefaultCurrentDayMediaPortalCountByStudentIds(ctx context.Context, studentIDs []string) error {
	if len(studentIDs) == 0 {
		return nil
	}

	currentDate := helper.GetCurrentDateStringWithTimezone()

	// TotalPlaceholder = 5 cho tung hoc sinh (chi append ngay hien tai neu chua co)
	for _, studentID := range studentIDs {
		if studentID == "" {
			continue
		}

		cacheKey := keys.GetCountCacheKey(studentID)

		var studentPortalCounts []media.CountMediaStudentPortal
		if err := s.cache.Get(ctx, cacheKey, &studentPortalCounts); err != nil {
			return err
		}

		isCurrentDayExists := false
		for _, studentPortalCount := range studentPortalCounts {
			if studentPortalCount.Date == currentDate {
				isCurrentDayExists = true
				break
			}
		}
		if isCurrentDayExists {
			continue
		}

		studentPortalCounts = append(studentPortalCounts, media.CountMediaStudentPortal{
			Date:             currentDate,
			TotalPlaceholder: 5,
			TotalResource:    0,
			TotalOutput:      0,
		})

		if err := cached.SetCache(s.cache, ctx, cacheKey, studentPortalCounts); err != nil {
			return err
		}
	}

	return nil
}

func (s *cachingCountService) InvalidateMediaPortalCountByStudentID(ctx context.Context, studentID string) error {
	if studentID == "" {
		return nil
	}
	return cached.DeleteCache(s.cache, ctx, keys.GetCountCacheKey(studentID))
}
