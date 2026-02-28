package webDevAccess

import "github.com/yijiacode188/wxSDK/service/handler/base"

// WebDevAccess 网页开发 网页授权
type WebDevAccess struct {
	*base.Base
}

func NewWebDevAccess(base *base.Base) (*WebDevAccess, error) {
	client := &WebDevAccess{
		Base: base,
	}

	return client, nil
}
