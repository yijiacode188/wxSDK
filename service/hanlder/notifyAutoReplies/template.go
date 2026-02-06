package notifyAutoReplies

import "github.com/yijiacode188/wxSDK/service/hanlder/base"

// NotifyAutoReplies 基础消息与订阅通知 自动回复
type NotifyAutoReplies struct {
	*base.Base
}

func NewNotifyAutoReplies(base *base.Base) (*NotifyAutoReplies, error) {
	client := &NotifyAutoReplies{
		Base: base,
	}

	return client, nil
}
