package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetTermCacheKey(termID string) string {
	return string(helper.TermServicePrefix) + "term:" + termID
}
