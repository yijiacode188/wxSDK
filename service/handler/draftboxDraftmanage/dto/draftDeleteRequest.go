package dto

import "encoding/json"

type DraftDeleteRequest struct {
	MediaId string `json:"media_id" comment:"要删除的草稿的media_id"`
}

func (p *DraftDeleteRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
