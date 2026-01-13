package keys

import (
	"strconv"

	"github.com/hung-senbox/senbox-cache-service/helper"
)

func GetLanguageByIdCacheKey(languageID uint) string {
	return string(helper.CommonServicePrefix) + "language_by_id:" + strconv.FormatUint(uint64(languageID), 10)
}
