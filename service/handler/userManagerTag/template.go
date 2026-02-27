package userManagerTag

import "github.com/yijiacode188/wxSDK/service/handler/base"

// UserManagerTag 用户管理 标签管理
type UserManagerTag struct {
	*base.Base
}

func NewUserManagerTag(base *base.Base) (*UserManagerTag, error) {
	client := &UserManagerTag{
		Base: base,
	}

	return client, nil
}
