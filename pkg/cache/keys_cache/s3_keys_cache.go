package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetS3UrlByKeyCacheKey(key string) string {
	return string(helper.S3ServicePrefix) + "s3-url-by-key:" + key
}
