package helper

type CachePrefix string

const (
	MainServicePrefixCache CachePrefix = "main-service:"
	ProfileServicePrefix   CachePrefix = "profile-service:"
	CommonServicePrefix    CachePrefix = "common-service:"
	LocationServicePrefix  CachePrefix = "location-service:"
	ColortimeServicePrefix CachePrefix = "colortime-service:"
	ReportServicePrefix    CachePrefix = "report-service:"
	QRServicePrefix        CachePrefix = "qr-service:"
	InventoryServicePrefix CachePrefix = "inventory-service:"
	HolidayServicePrefix   CachePrefix = "holiday-service:"
	TermServicePrefix      CachePrefix = "term-service:"
	S3ServicePrefix        CachePrefix = "s3-service:"
	CountServicePrefix     CachePrefix = "count-service:"
	MediaPrefix            CachePrefix = "media:"
)

type MediaCountSubType string

const (
	MediaCountSubTypePortalByStudentID MediaCountSubType = "portal_by_student_id"
)
