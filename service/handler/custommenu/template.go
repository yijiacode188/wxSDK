package custommenu

import (
	"github.com/yijiacode188/wxSDK/service/handler/base"
)

type CustomMenu struct {
	*base.Base
}

func NewCustomMenu(base *base.Base) (*CustomMenu, error) {
	client := &CustomMenu{
		Base: base,
	}
	return client, nil
}
