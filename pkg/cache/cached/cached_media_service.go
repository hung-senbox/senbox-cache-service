package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/media"
)

type CachedMediaService interface {
	GetTopicByID(ctx context.Context, topicID string) (*media.Topic, error)
	GetTopics(ctx context.Context, topicIDs []string) (map[string]*media.Topic, error)
}

type cachedMediaService struct {
	cache     *cache.RedisCache
	s3Service CachedS3Service
}

func NewCachedMediaService(cache *cache.RedisCache) CachedMediaService {
	return &cachedMediaService{
		cache:     cache,
		s3Service: NewCachedS3Service(cache),
	}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedMediaService) GetTopicByID(ctx context.Context, topicID string) (*media.Topic, error) {
	if topicID == "" {
		return nil, nil
	}

	var topic *media.Topic
	if err := c.cache.Get(ctx, keys.GetTopicByIDCacheKey(topicID), &topic); err != nil {
		return nil, err
	}
	if topic == nil {
		return nil, nil
	}

	if topic.MainImageKey != "" {
		url, err := c.s3Service.GetS3UrlByKey(ctx, topic.MainImageKey)
		if err != nil {
			return nil, err
		}
		topic.MainImageUrl = url
	}

	return topic, nil
}

func (c *cachedMediaService) GetTopics(ctx context.Context, topicIDs []string) (map[string]*media.Topic, error) {
	if len(topicIDs) == 0 {
		return nil, nil
	}

	result := make(map[string]*media.Topic, len(topicIDs))
	for _, topicID := range topicIDs {
		if topicID == "" {
			continue
		}

		topic, err := c.GetTopicByID(ctx, topicID)
		if err != nil {
			return nil, err
		}
		if topic == nil {
			continue
		}

		result[topicID] = topic
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
