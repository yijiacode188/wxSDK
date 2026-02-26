package dto

import "encoding/json"

type FreePublishSubmitRequest struct {
	MediaId string `json:"media_id" comment:"要发布的草稿的media_id"`
}

func (p *FreePublishSubmitRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
