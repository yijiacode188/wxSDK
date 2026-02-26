package materialPermanent

import "github.com/yijiacode188/wxSDK/service/handler/base"

// MaterialPermanent 素材管理 永久素材
type MaterialPermanent struct {
	*base.Base
}

func NewMaterialPermanent(base *base.Base) (*MaterialPermanent, error) {
	client := &MaterialPermanent{
		Base: base,
	}
	return client, nil
}
