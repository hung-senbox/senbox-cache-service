package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

// Check In Cache Key
func UserCheckInCacheKey(userID string) string {
	return string(helper.HolidayServicePrefix) + "user-check-in:" + userID
}

func StudentCheckInCacheKey(studentID string) string {
	return string(helper.HolidayServicePrefix) + "student-check-in:" + studentID
}

// Check Out Cache Key

func UserCheckOutCacheKey(userID string) string {
	return string(helper.HolidayServicePrefix) + "user-check-out:" + userID
}

func StudentCheckOutCacheKey(studentID string) string {
	return string(helper.HolidayServicePrefix) + "student-check-out:" + studentID
}
