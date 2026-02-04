package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo/base"

type GetStableAccessTokenResponse struct {
	base.Base
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
