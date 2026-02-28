package event

import (
	"github.com/yijiacode188/wxSDK/service/handler/event/dto"
)

// Auth 微信事件消息验证消息
func (wx *Event) Auth(body *dto.AuthRequest) bool {
	return validateAuthPlainText(wx.Token, body.Timestamp, body.Nonce, body.Signature)
}
