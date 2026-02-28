package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type SnsAccessTokenResponse struct {
	vo.Base
	AccessToken    string `json:"access_token" comment:"网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同"`
	ExpiresIn      int64  `json:"expires_in" comment:"access_token接口调用凭证超时时间，单位（秒）"`
	RefreshToken   string `json:"refresh_token" comment:"用户刷新access_token"`
	OpenID         string `json:"openid" comment:"用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID"`
	UnIoId         string `json:"unionid" comment:"用户统一标识（针对一个微信开放平台账号下的应用，同一用户的 unionid 是唯一的），只有当scope为\"snsapi_userinfo\"时返回"`
	IsSnapshotUser int    `json:"is_snapshotuser" comment:"是否为快照页模式虚拟账号，只有当用户是快照页模式虚拟账号时返回，值为1"`
}
