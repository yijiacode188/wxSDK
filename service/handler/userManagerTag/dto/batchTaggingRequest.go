package dto

import "encoding/json"

type BatchTaggingRequest struct {
	OpenIdList []string `json:"openid_list" comment:"粉丝openid列表，最多50个"`
	TagId      int64    `json:"tagid" comment:"标签id"`
}

func (p *BatchTaggingRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
