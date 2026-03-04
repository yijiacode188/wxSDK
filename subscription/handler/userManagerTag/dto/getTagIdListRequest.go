package dto

import "encoding/json"

type GetTagIdListRequest struct {
	OpenId string `json:"openid" comment:"用户 openid"`
}

func (p *GetTagIdListRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
