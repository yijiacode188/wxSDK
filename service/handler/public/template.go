package public

import "github.com/yijiacode188/wxSDK/service/handler/base"

// Public 发布能力
type Public struct {
	*base.Base
}

func NewPublic(base *base.Base) (*Public, error) {
	client := &Public{
		Base: base,
	}

	return client, nil
}
