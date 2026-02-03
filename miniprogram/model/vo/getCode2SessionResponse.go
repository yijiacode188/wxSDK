package vo

import "wxSDK/miniprogram/model/vo/base"

type GetCode2SessionResponse struct {
	base.Base
	SessionKey string `json:"session_key" comment:"会话秘钥"`
	UnionId    string `json:"unionid" comment:"用户在开放平台的唯一标识符"`
	OpenId     string `json:"openid" comment:"用户唯一标识"`
}
