package dto

import "encoding/json"

type BatchBlackListRequest struct {
	OpenIdList []string `json:"openid_list" comment:"需要拉黑的openid列表"`
}

func (p *BatchBlackListRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
