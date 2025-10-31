package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func UserCacheKey(userID string) string {
	return string(helper.MainServicePrefixCache) + "user:" + userID
}

func StudentCacheKey(studentID string) string {
	return string(helper.MainServicePrefixCache) + "student:" + studentID
}

func TeacherCacheKey(teacherID string) string {
	return string(helper.MainServicePrefixCache) + "teacher:" + teacherID
}

func StaffCacheKey(staffID string) string {
	return string(helper.MainServicePrefixCache) + "staff:" + staffID
}

func ParentCacheKey(parentID string) string {
	return string(helper.MainServicePrefixCache) + "parent:" + parentID
}

func ChildCacheKey(childID string) string {
	return string(helper.MainServicePrefixCache) + "child:" + childID
}

func TeacherByUserAndOrgCacheKey(userID, orgID string) string {
	return string(helper.MainServicePrefixCache) + "teacher-by-user-org:" + userID + ":" + orgID
}

func StaffByUserAndOrgCacheKey(userID, orgID string) string {
	return string(helper.MainServicePrefixCache) + "staff-by-user-org:" + userID + ":" + orgID
}

func UserByTeacherCacheKey(teacherID string) string {
	return string(helper.MainServicePrefixCache) + "user-by-teacher:" + teacherID
}

func ParentByUserCacheKey(userID string) string {
	return string(helper.MainServicePrefixCache) + "parent-by-user:" + userID
}
