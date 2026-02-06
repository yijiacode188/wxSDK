package dto

import "encoding/json"

type ClearQuotaRequest struct {
	AppId string `json:"appid" comment:"要被清空的账号的appid"`
}

func (p *ClearQuotaRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
