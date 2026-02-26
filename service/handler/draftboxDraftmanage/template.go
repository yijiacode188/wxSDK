package draftboxDraftmanage

import "github.com/yijiacode188/wxSDK/service/handler/base"

// DraftBoxDraftManage 草稿管理和商品卡片 草稿管理
type DraftBoxDraftManage struct {
	*base.Base
}

func NewDraftBoxDraftManage(base *base.Base) (*DraftBoxDraftManage, error) {
	client := &DraftBoxDraftManage{
		Base: base,
	}
	return client, nil
}
