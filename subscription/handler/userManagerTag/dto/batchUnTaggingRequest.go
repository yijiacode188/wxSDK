package dto

import "encoding/json"

type BatchUnTaggingRequest struct {
	OpenIdList []string `json:"openid_list" comment:"粉丝openid列表"`
	TagId      int64    `json:"tagid" comment:"标签id"`
}

func (p *BatchUnTaggingRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
