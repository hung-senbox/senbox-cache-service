package helper

type CachePrefix string

const (
	MainServicePrefixCache CachePrefix = "main-service:"
	ProfileServicePrefix   CachePrefix = "profile-service:"
	CommonServicePrefix    CachePrefix = "common-service:"
	LocationServicePrefix  CachePrefix = "location-service:"
	ReportServicePrefix    CachePrefix = "report-service:"
	QRServicePrefix        CachePrefix = "qr-service:"
	InventoryServicePrefix CachePrefix = "inventory-service:"
)
