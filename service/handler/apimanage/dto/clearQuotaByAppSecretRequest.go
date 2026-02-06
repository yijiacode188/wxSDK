package dto

import "encoding/json"

type ClearQuotaByAppSecretRequest struct {
	AppID     string `json:"appid" comment:"要被清空的账号的appid"`
	AppSecret string `json:"appsecret" comment:"唯一凭证密钥，即 AppSecret"`
}

func (p *ClearQuotaByAppSecretRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
