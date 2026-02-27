package dto

import "encoding/json"

type UpdateRemarkRequest struct {
	OpenId string `json:"openid" comment:"openId"`
	Remark string `json:"remark" comment:"备注名"`
}

func (p *UpdateRemarkRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
