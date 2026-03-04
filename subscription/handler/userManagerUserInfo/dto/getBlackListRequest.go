package dto

import "encoding/json"

type GetBlackListRequest struct {
	BeginOpenId string `json:"begin_openid" comment:"起始OpenID，为空时从开头拉取"`
}

func (p *GetBlackListRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
