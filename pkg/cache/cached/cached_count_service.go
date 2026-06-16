package cached

import (
	"context"
	"fmt"
	"time"

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

func (c *cachedCountService) GetMediaPortalCountByStudentIdsCurrentDay(ctx context.Context, studentIDs []string) (map[string][]media.CountMediaStudentPortal, error) {
	if len(studentIDs) == 0 {
		return nil, nil
	}

	currentDate := time.Now().Format("2006-01-02")
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

		currentDayCounts := make([]media.CountMediaStudentPortal, 0, 1)
		for _, studentPortalCount := range studentPortalCounts {
			if studentPortalCount.Date != currentDate {
				continue
			}
			currentDayCounts = append(currentDayCounts, studentPortalCount)
		}

		if len(currentDayCounts) == 0 {
			continue
		}

		result[studentID] = currentDayCounts
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func (c *cachedCountService) GetMediaPortalCountByStudentIdsRangeData(ctx context.Context, studentIDs []string, startDate, endDate string) (map[string][]media.CountMediaStudentPortal, error) {
	if len(studentIDs) == 0 {
		return nil, nil
	}
	if startDate == "" || endDate == "" {
		return nil, nil
	}

	const layout = "2006-01-02"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid startDate format, expected yyyy-mm-dd: %w", err)
	}
	end, err := time.Parse(layout, endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid endDate format, expected yyyy-mm-dd: %w", err)
	}
	if end.Before(start) {
		return nil, fmt.Errorf("endDate must be greater than or equal to startDate")
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

		rangeCounts := make([]media.CountMediaStudentPortal, 0, len(studentPortalCounts))
		for _, studentPortalCount := range studentPortalCounts {
			countDate, err := time.Parse(layout, studentPortalCount.Date)
			if err != nil {
				continue
			}
			if countDate.Before(start) || countDate.After(end) {
				continue
			}
			rangeCounts = append(rangeCounts, studentPortalCount)
		}

		if len(rangeCounts) == 0 {
			continue
		}

		result[studentID] = rangeCounts
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
