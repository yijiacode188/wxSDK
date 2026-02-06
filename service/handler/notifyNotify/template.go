package notifyNotify

import "github.com/yijiacode188/wxSDK/service/handler/base"

// NotifyNotify 基础消息与订阅通知 订阅通知
type NotifyNotify struct {
	*base.Base
}

func NewNotifyNotify(base *base.Base) (*NotifyNotify, error) {
	client := &NotifyNotify{
		Base: base,
	}

	return client, nil
}
