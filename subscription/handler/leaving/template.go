package leaving

import "github.com/yijiacode188/wxSDK/subscription/handler/base"

// Leaving 留言管理
type Leaving struct {
	*base.Base
}

func NewLeaving(base *base.Base) (*Leaving, error) {
	client := &Leaving{
		Base: base,
	}
	return client, nil
}
