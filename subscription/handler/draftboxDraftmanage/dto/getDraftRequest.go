package dto

import "encoding/json"

type GetDraftRequest struct {
	MediaId string `json:"media_id" comment:"要获取的草稿的media_id"`
}

func (p *GetDraftRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
