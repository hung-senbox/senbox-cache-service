package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetTopicByIDCacheKey(topicID string) string {
	return string(helper.MediaPrefix) + "topic:" + topicID
}
