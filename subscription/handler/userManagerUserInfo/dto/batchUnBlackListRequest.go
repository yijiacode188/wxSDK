package dto

import "encoding/json"

type BatchUnBlackListRequest struct {
	OpenIdList []string `json:"openid_list" comment:"需要取消拉黑的openid列表"`
}

func (p *BatchUnBlackListRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
