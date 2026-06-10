package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

// media
func GetCountCacheKey(studentID string) string {
	return string(helper.CountServicePrefix) + string(helper.MediaPrefix) + string(helper.MediaCountSubTypePortalByStudentID) + studentID
}
