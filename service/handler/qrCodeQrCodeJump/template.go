package qrCodeQrCodeJump

import "github.com/yijiacode188/wxSDK/service/handler/base"

// QrCodeQrCodeJump 服务号二维码 扫二维码打开小程序
type QrCodeQrCodeJump struct {
	*base.Base
}

func NewQrCodeQrCodeJump(base *base.Base) (*QrCodeQrCodeJump, error) {
	client := &QrCodeQrCodeJump{
		Base: base,
	}
	return client, nil
}
