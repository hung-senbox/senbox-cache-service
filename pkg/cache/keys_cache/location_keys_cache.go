package keys

import (
	"strconv"

	"github.com/hung-senbox/senbox-cache-service/helper"
)

func GetLocationByIdCacheKey(locationID uint) string {
	return string(helper.LocationServicePrefix) + "location_by_id:" + strconv.FormatUint(uint64(locationID), 10)
}
