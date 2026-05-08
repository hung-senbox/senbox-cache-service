package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetColorTimeSessionIDCacheKey(colorTimeSessionID string) string {
	return string(helper.ColortimeServicePrefix) + "colortime_session_id:" + colorTimeSessionID
}
