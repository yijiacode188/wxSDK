package openPocAi

import "github.com/yijiacode188/wxSDK/service/handler/base"

// OpenPocAi 智能接口 AI开放接口
type OpenPocAi struct {
	*base.Base
}

func NewOpenPocAi(base *base.Base) (*OpenPocAi, error) {
	client := &OpenPocAi{
		Base: base,
	}

	return client, nil
}
