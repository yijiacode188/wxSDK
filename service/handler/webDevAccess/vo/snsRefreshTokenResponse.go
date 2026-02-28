package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type SnsRefreshTokenResponse struct {
	vo.Base
	AccessToken  string `json:"access_token" comment:"网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同"`
	ExpiresIn    int64  `json:"expires_in" comment:"access_token接口调用凭证超时时间，单位（秒）"`
	RefreshToken string `json:"refresh_token" comment:"用户刷新access_token"`
	OpenID       string `json:"openid" comment:"用户唯一标识"`
	Scope        string `json:"scope" comment:"用户授权的作用域，使用逗号（,）分隔"`
}
