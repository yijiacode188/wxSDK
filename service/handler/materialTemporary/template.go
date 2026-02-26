package materialTemporary

import "github.com/yijiacode188/wxSDK/service/handler/base"

// MaterialTemporary 素材管理 临时素材
type MaterialTemporary struct {
	*base.Base
}

func NewMaterialTemporary(base *base.Base) (*MaterialTemporary, error) {
	client := &MaterialTemporary{
		Base: base,
	}
	return client, nil
}
