package dto

import "encoding/json"

type QrCodeJumpPublishRequest struct {
	Prefix string `json:"prefix" comment:"二维码规则。如果是服务号，则是服务号的带参的二维码url。"`
}

func (p *QrCodeJumpPublishRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
