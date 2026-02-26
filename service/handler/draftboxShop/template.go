package draftboxShop

import "github.com/yijiacode188/wxSDK/service/handler/base"

// DraftBoxShop 草稿管理和商品卡片 商品卡片
type DraftBoxShop struct {
	*base.Base
}

func NewDraftBoxShop(base *base.Base) (*DraftBoxShop, error) {
	client := &DraftBoxShop{
		Base: base,
	}
	return client, nil
}
