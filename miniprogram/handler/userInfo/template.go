package userInfo

import (
	"github.com/yijiacode188/wxSDK/miniprogram/handler/mpAccessToken"
)

type UserInfo struct {
	*mpAccessToken.MPAccessToken
}

func NewUserInfo(mp *mpAccessToken.MPAccessToken) (*UserInfo, error) {
	client := &UserInfo{
		MPAccessToken: mp,
	}
	return client, nil
}
