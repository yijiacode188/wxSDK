package vo

import "wxSDK/miniprogram/model/vo/base"

type GetAccessTokenResponse struct {
	base.Base
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
