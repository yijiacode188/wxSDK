package dto

import "encoding/json"

type DelMaterialRequest struct {
	MediaId string `json:"media_id" comment:"要删除的素材media_id"`
}

func (p *DelMaterialRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
