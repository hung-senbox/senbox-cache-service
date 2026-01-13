package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetLocationByIdCacheKey(locationID string) string {
	return string(helper.LocationServicePrefix) + "location_by_id:" + locationID
}
