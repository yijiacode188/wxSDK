package dto

import (
	"github.com/yijiacode188/wxSDK/service/handler/event/vo"
)

// EventMessageBody 消息主体
type EventMessageBody struct {
	vo.EventMessageResponse
	Encrypt string `json:"Encrypt"`
}
