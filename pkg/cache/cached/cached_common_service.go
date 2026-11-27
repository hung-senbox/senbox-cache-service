package cached

import (
	"context"

	"github.com/hung-senbox/senbox-cache-service/pkg/cache"
	keys "github.com/hung-senbox/senbox-cache-service/pkg/cache/keys_cache"
	"github.com/hung-senbox/senbox-cache-service/pkg/model/common"
)

type CachedCommonGateway interface {
	GetOrganizationSettingByOrganizationID(ctx context.Context, organizationID string) (*common.OrganizationSetting, error)
	GetOrganizationSettingByOrganizationIds(ctx context.Context, organizationIds []string) (map[string]*common.OrganizationSetting, error)

	// ======================== Language Cache ========================
	GetLanguageById(ctx context.Context, languageID uint) (map[string]interface{}, error)

	// ======================== Occupation Cache ========================
	GetOccupationByIdAndLanguageId(ctx context.Context, occupationID string, languageID uint) (map[string]interface{}, error)
}

type cachedCommonService struct {
	cache *cache.RedisCache
}

func NewCachedCommonGateway(cache *cache.RedisCache) CachedCommonGateway {
	return &cachedCommonService{
		cache: cache,
	}
}

// ========================
// === GET CACHE ===
// ========================

func (c *cachedCommonService) GetLanguageById(ctx context.Context, languageID uint) (map[string]interface{}, error) {
	if languageID == 0 {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetLanguageByIdCacheKey(languageID))
}

func (c *cachedCommonService) GetOccupationByIdAndLanguageId(ctx context.Context, occupationID string, languageID uint) (map[string]interface{}, error) {
	if occupationID == "" || languageID == 0 {
		return nil, nil
	}
	return getCache(c.cache, ctx, keys.GetOccupationByIdAndLanguageIdCacheKey(occupationID, languageID))
}

func (c *cachedCommonService) GetOrganizationSettingByOrganizationID(ctx context.Context, organizationID string) (*common.OrganizationSetting, error) {
	if organizationID == "" {
		return nil, nil
	}

	var setting *common.OrganizationSetting
	if err := c.cache.Get(ctx, keys.GetOrganizationSettingByOrganizationIDCacheKey(organizationID), &setting); err != nil {
		return nil, err
	}
	return setting, nil
}

func (c *cachedCommonService) GetOrganizationSettingByOrganizationIds(ctx context.Context, organizationIds []string) (map[string]*common.OrganizationSetting, error) {
	settings := make(map[string]*common.OrganizationSetting)
	seen := make(map[string]struct{}, len(organizationIds))
	for _, organizationID := range organizationIds {
		if organizationID == "" {
			continue
		}
		if _, exists := seen[organizationID]; exists {
			continue
		}
		seen[organizationID] = struct{}{}

		setting, err := c.GetOrganizationSettingByOrganizationID(ctx, organizationID)
		if err != nil {
			return nil, err
		}
		if setting != nil {
			settings[organizationID] = setting
		}
	}
	return settings, nil
}
