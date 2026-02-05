package userLogin

import (
	"github.com/yijiacode188/wxSDK/miniprogram/handler/mpAccessToken"
)

type UserLogin struct {
	*mpAccessToken.MPAccessToken
}

func NewUserLogin(mp *mpAccessToken.MPAccessToken) (*UserLogin, error) {
	client := &UserLogin{
		MPAccessToken: mp,
	}
	return client, nil
}
