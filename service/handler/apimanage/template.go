package apimanage

import (
	"github.com/yijiacode188/wxSDK/service/handler/base"
)

type ApiManager struct {
	*base.Base
}

func NewApiManager(base *base.Base) (*ApiManager, error) {
	client := &ApiManager{
		Base: base,
	}
	return client, nil
}
