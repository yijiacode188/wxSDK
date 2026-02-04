package dto

import "encoding/json"

type GetStableAccessTokenRequest struct {
	GrantType    string `json:"grant_type" comment:"填写 client_credential"`
	AppID        string `json:"appid" comment:"账号的唯一凭证，即 AppID"`
	Secret       string `json:"secret" comment:"唯一凭证密钥，即 AppSecret"`
	ForceRefresh bool   `json:"force_refresh" comment:"默认使用 false。1. force_refresh = false 时为普通调用模式，access_token 有效期内重复调用该接口不会更新 access_token；2. 当force_refresh = true 时为强制刷新模式，会导致上次获取的 access_token 失效，并返回新的 access_token"`
}

func (p *GetStableAccessTokenRequest) ToByte() []byte {
	p.GrantType = "client_credential"
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
