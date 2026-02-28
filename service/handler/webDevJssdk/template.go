package webDevJssdk

import "github.com/yijiacode188/wxSDK/service/handler/base"

// WebDevJsSdk 网页开发 JSSDK
type WebDevJsSdk struct {
	*base.Base
}

func NewWebDevJsSdk(base *base.Base) (*WebDevJsSdk, error) {
	client := &WebDevJsSdk{
		Base: base,
	}

	return client, nil
}
