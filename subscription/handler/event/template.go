package event

import "github.com/yijiacode188/wxSDK/subscription/handler/base"

// Event 事件通知
type Event struct {
	*base.Base
}

func NewEvent(base *base.Base) (*Event, error) {
	client := &Event{
		Base: base,
	}
	return client, nil
}
