package helper

type CachePrefix string

const (
	MainServicePrefixCache CachePrefix = "main-service:"
	ProfileServicePrefix   CachePrefix = "profile-service:"
	CommonServicePrefix    CachePrefix = "common-service:"
	LocationServicePrefix  CachePrefix = "location-service:"
)
