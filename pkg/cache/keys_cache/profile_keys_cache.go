package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func UserCodeCacheKey(userID string) string {
	return string(helper.ProfileServicePrefix) + "user_code:" + userID
}

func StudentCodeCacheKey(studentID string) string {
	return string(helper.ProfileServicePrefix) + "student_code:" + studentID
}

func TeacherCodeCacheKey(teacherID string) string {
	return string(helper.ProfileServicePrefix) + "teacher_code:" + teacherID
}

func StaffCodeCacheKey(staffID string) string {
	return string(helper.ProfileServicePrefix) + "staff_code:" + staffID
}

func ParentCodeCacheKey(parentID string) string {
	return string(helper.ProfileServicePrefix) + "parent_code:" + parentID
}

func ChildCodeCacheKey(childID string) string {
	return string(helper.ProfileServicePrefix) + "child_code:" + childID
}

func DeviceCodeCacheKey(deviceID string) string {
	return string(helper.ProfileServicePrefix) + "device_code:" + deviceID
}
