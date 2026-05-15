package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetTermCacheKey(termID string) string {
	return string(helper.TermServicePrefix) + "term:" + termID
}

func GetCurrentTermByOrganizationIDCacheKey(organizationID string) string {
	return string(helper.TermServicePrefix) + "current-term-by-organization-id:" + organizationID
}
