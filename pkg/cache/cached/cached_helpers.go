package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
)

// ========================
// === HELPER METHODS ===
// ========================

// getCache retrieves a map[string]interface{} from cache using the provided cache key
func getCache(
	cache *cache.RedisCache,
	ctx context.Context,
	cacheKey string,
) (map[string]interface{}, error) {
	if cacheKey == "" {
		return nil, nil
	}

	var result map[string]interface{}
	if err := cache.Get(ctx, cacheKey, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

// getCache4ParentReportLanguages retrieves an array of maps from cache for parent report languages
func getCache4ParentReportLanguages(
	cache *cache.RedisCache,
	ctx context.Context,
	cacheKey string,
) ([]map[string]interface{}, error) {
	if cacheKey == "" {
		return nil, nil
	}

	var result []map[string]interface{}
	if err := cache.Get(ctx, cacheKey, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func getCacheArray(
	cache *cache.RedisCache,
	ctx context.Context,
	cacheKey string,
) ([]map[string]interface{}, error) {
	if cacheKey == "" {
		return nil, nil
	}
	var result []map[string]interface{}
	if err := cache.Get(ctx, cacheKey, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func getCacheString(
	cache *cache.RedisCache,
	ctx context.Context,
	cacheKey string,
) (string, error) {
	if cacheKey == "" {
		return "", nil
	}

	var result string
	if err := cache.Get(ctx, cacheKey, &result); err != nil {
		return "", err
	}

	if result == "" {
		return "", nil
	}

	return result, nil
}

func getCacheFloat(
	cache *cache.RedisCache,
	ctx context.Context,
	cacheKey string,
) (float32, error) {
	if cacheKey == "" {
		return 0, nil
	}

	var result float32
	if err := cache.Get(ctx, cacheKey, &result); err != nil {
		return 0, err
	}

	if result == 0 {
		return 0, nil
	}

	return result, nil
}
