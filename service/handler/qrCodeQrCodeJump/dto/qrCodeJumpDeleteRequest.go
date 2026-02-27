package dto

import "encoding/json"

type QrCodeJumpDeleteRequest struct {
	Prefix string `json:"prefix" comment:"二维码规则"`
	AppID  string `json:"appid" comment:"小程序appid"`
}

func (p *QrCodeJumpDeleteRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
