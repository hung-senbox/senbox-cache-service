package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

// Check In Cache Key
func StudentCheckInCacheKey(studentID string) string {
	return string(helper.HolidayServicePrefix) + "student-check-in:" + studentID
}

func TeacherCheckInCacheKey(teacherID string) string {
	return string(helper.HolidayServicePrefix) + "teacher-check-in:" + teacherID
}

func StaffCheckInCacheKey(staffID string) string {
	return string(helper.HolidayServicePrefix) + "staff-check-in:" + staffID
}

// Check Out Cache Key
func StudentCheckOutCacheKey(studentID string) string {
	return string(helper.HolidayServicePrefix) + "student-check-out:" + studentID
}

func TeacherCheckOutCacheKey(teacherID string) string {
	return string(helper.HolidayServicePrefix) + "teacher-check-out:" + teacherID
}

func StaffCheckOutCacheKey(staffID string) string {
	return string(helper.HolidayServicePrefix) + "staff-check-out:" + staffID
}
