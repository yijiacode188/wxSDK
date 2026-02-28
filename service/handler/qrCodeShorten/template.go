package qrCodeShorten

import "github.com/yijiacode188/wxSDK/service/handler/base"

// QrCodeShorten 服务号二维码 长信息与短链
type QrCodeShorten struct {
	*base.Base
}

func NewQrCodeShorten(base *base.Base) (*QrCodeShorten, error) {
	client := &QrCodeShorten{
		Base: base,
	}
	return client, nil
}
