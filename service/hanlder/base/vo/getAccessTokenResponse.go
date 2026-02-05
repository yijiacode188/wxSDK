package vo

import (
	"github.com/yijiacode188/wxSDK/service/model/vo"
)

type GetAccessTokenResponse struct {
	vo.Base
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
