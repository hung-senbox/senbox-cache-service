package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetShelfSlotByIdCacheKey(shelfSlotID string) string {
	return string(helper.InventoryServicePrefix) + "shelf_slot_by_id:" + shelfSlotID
}
