package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
)

type CachedS3Service interface {
	GetS3UrlByKey(ctx context.Context, key string) (string, error)
	GetS3UrlByKeys(ctx context.Context, objectKeys []string) (map[string]string, error)
}

type cachedS3Service struct {
	cache *cache.RedisCache
}

func NewCachedS3Service(cache *cache.RedisCache) CachedS3Service {
	return &cachedS3Service{cache: cache}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedS3Service) GetS3UrlByKey(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", nil
	}
	return getCacheString(c.cache, ctx, keys.GetS3UrlByKeyCacheKey(key))
}

func (c *cachedS3Service) GetS3UrlByKeys(ctx context.Context, objectKeys []string) (map[string]string, error) {
	if len(objectKeys) == 0 {
		return nil, nil
	}

	result := make(map[string]string, len(objectKeys))
	for _, objectKey := range objectKeys {
		if objectKey == "" {
			continue
		}

		url, err := c.GetS3UrlByKey(ctx, objectKey)
		if err != nil {
			return nil, err
		}
		if url == "" {
			continue
		}

		result[objectKey] = url
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
