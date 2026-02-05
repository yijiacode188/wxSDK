package notifyMessage

import "github.com/yijiacode188/wxSDK/service/hanlder/base"

// NotifyMessage 基础消息与订阅通知
type NotifyMessage struct {
	*base.Base
}

func NewNotifyMessage(base *base.Base) (*NotifyMessage, error) {
	client := &NotifyMessage{
		Base: base,
	}

	return client, nil
}
