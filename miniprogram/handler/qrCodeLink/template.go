package qrCodeLink

import (
	"github.com/yijiacode188/wxSDK/miniprogram/handler/mpAccessToken"
)

type QRCodeLink struct {
	*mpAccessToken.MPAccessToken
}

func NewQRCodeLink(mp *mpAccessToken.MPAccessToken) (*QRCodeLink, error) {
	client := &QRCodeLink{
		MPAccessToken: mp,
	}
	return client, nil
}
