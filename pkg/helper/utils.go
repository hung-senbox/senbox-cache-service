package helper

import "time"

func GetCurrentDateStringWithTimezone() string {
	loc := time.FixedZone("GMT+7", 7*60*60)
	return time.Now().In(loc).Format("2006-01-02")
}
