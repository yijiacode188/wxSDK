package qrCodeQrcodes

import "github.com/yijiacode188/wxSDK/service/handler/base"

// QrCodeQrCodes 服务号二维码 带参二维码
type QrCodeQrCodes struct {
	*base.Base
}

func NewQrCodeQrCodes(base *base.Base) (*QrCodeQrCodes, error) {
	client := &QrCodeQrCodes{
		Base: base,
	}
	return client, nil
}
