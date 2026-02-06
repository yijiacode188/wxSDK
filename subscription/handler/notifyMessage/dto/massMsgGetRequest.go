package dto

import "encoding/json"

type MassMsgGetRequest struct {
	MsgId string `json:"msg_id" comment:"群发消息后返回的消息id"`
}

func (p *MassMsgGetRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
