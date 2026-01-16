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

func OrganizationCodeCacheKey(organizationID string) string {
	return string(helper.ProfileServicePrefix) + "organization_code:" + organizationID
}

func ChildEnrollmentCodeCacheKey(childID string) string {
	return string(helper.ProfileServicePrefix) + "child_enrollment:" + childID
}

// language setting cache key
func StudentLanguageSettingCacheKey(studentID string) string {
	return string(helper.ProfileServicePrefix) + "student_language_setting:" + studentID
}

func TeacherLanguageSettingCacheKey(teacherID string) string {
	return string(helper.ProfileServicePrefix) + "teacher_language_setting:" + teacherID
}

func StaffLanguageSettingCacheKey(staffID string) string {
	return string(helper.ProfileServicePrefix) + "staff_language_setting:" + staffID
}

func ParentLanguageSettingCacheKey(parentID string) string {
	return string(helper.ProfileServicePrefix) + "parent_language_setting:" + parentID
}

func ChildLanguageSettingCacheKey(childID string) string {
	return string(helper.ProfileServicePrefix) + "child_language_setting:" + childID
}

// Blocked User Cache Key
func BlockedUserCacheKey(userID string) string {
	return string(helper.ProfileServicePrefix) + "blocked_user:" + userID
}

func BlockedStudentCacheKey(studentID string) string {
	return string(helper.ProfileServicePrefix) + "blocked_student:" + studentID
}

func BlockedTeacherCacheKey(teacherID string) string {
	return string(helper.ProfileServicePrefix) + "blocked_teacher:" + teacherID
}

func BlockedStaffCacheKey(staffID string) string {
	return string(helper.ProfileServicePrefix) + "blocked_staff:" + staffID
}

func BlockedParentCacheKey(parentID string) string {
	return string(helper.ProfileServicePrefix) + "blocked_parent:" + parentID
}

func BlockedChildCacheKey(childID string) string {
	return string(helper.ProfileServicePrefix) + "blocked_child:" + childID
}
