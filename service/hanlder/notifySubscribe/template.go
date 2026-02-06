package notifySubscribe

import "github.com/yijiacode188/wxSDK/service/hanlder/base"

// NotifySubscribe 基础消息与订阅通知 一次性订阅消息
type NotifySubscribe struct {
	*base.Base
}

func NewNotifyNotify(base *base.Base) (*NotifySubscribe, error) {
	client := &NotifySubscribe{
		Base: base,
	}

	return client, nil
}
