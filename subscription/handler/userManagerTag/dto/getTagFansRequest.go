package dto

import "encoding/json"

type GetTagFansRequest struct {
	TagId      string `json:"tag_id" comment:"标签 ID"`
	NextOpenId string `json:"next_openid" comment:"第一个拉取的OPENID，不填默认从头开始拉取"`
}

func (p *GetTagFansRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
