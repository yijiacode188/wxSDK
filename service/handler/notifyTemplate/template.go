package notifyTemplate

import "github.com/yijiacode188/wxSDK/service/handler/base"

// NotifyTemplate 基础消息与订阅通知 模版消息
type NotifyTemplate struct {
	*base.Base
}

func NewNotifyTemplate(base *base.Base) (*NotifyTemplate, error) {
	client := &NotifyTemplate{
		Base: base,
	}

	return client, nil
}
