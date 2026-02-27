package userManagerUserInfo

import "github.com/yijiacode188/wxSDK/service/handler/base"

// UserManagerUserInfo 用户管理 用户信息
type UserManagerUserInfo struct {
	*base.Base
}

func NewUserManagerUserInfo(base *base.Base) (*UserManagerUserInfo, error) {
	client := &UserManagerUserInfo{
		Base: base,
	}

	return client, nil
}
