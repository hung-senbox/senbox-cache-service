package caching

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/media"
)

type CachingMediaService interface {
	SetTopicByID(ctx context.Context, topicID string, topic *media.Topic) error
	SetTopics(ctx context.Context, topics []*media.Topic) error
}

type cachingMediaService struct {
	cache      *cache.RedisCache
	defaultTTL int
}

func NewCachingMediaService(redisCache *cache.RedisCache, ttlSeconds int) CachingMediaService {
	return &cachingMediaService{
		cache:      redisCache,
		defaultTTL: ttlSeconds,
	}
}

func (s *cachingMediaService) SetTopicByID(ctx context.Context, topicID string, topic *media.Topic) error {
	if topicID == "" || topic == nil {
		return nil
	}
	return s.cache.Set(ctx, keys.GetTopicByIDCacheKey(topicID), topic, 0)
}

func (s *cachingMediaService) SetTopics(ctx context.Context, topics []*media.Topic) error {
	if len(topics) == 0 {
		return nil
	}
	for _, topic := range topics {
		if topic == nil || topic.ID == "" {
			continue
		}
		if err := s.SetTopicByID(ctx, topic.ID, topic); err != nil {
			return err
		}
	}
	return nil
}
