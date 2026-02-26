package dto

import "encoding/json"

type GetMaterialRequest struct {
	MediaId string `json:"media_id" comment:"要获取的素材的media_id"`
}

func (p *GetMaterialRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
