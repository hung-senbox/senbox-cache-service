package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetQRCodeCacheKey(qrCodeID string) string {
	return string(helper.QRServicePrefix) + "qr_code:" + qrCodeID
}

func GetLibModeCacheKey(libModeID string) string {
	return string(helper.QRServicePrefix) + "lib_mode:" + libModeID
}
