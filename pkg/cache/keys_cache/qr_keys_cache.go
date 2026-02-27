package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func GetQRCodeCacheKey(qrCodeID string) string {
	return string(helper.QRServicePrefix) + "qr_code:" + qrCodeID
}
