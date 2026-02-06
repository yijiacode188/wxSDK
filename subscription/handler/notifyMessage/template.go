package notifyMessage

import "github.com/yijiacode188/wxSDK/subscription/handler/base"

// NotifyMessage 基础消息与订阅通知 群发消息
type NotifyMessage struct {
	*base.Base
}

func NewNotifyMessage(base *base.Base) (*NotifyMessage, error) {
	client := &NotifyMessage{
		Base: base,
	}

	return client, nil
}
