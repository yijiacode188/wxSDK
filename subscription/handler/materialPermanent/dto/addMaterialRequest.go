package dto

import "encoding/json"

type AddMaterialRequest struct {
	Title        string `json:"title" comment:"视频素材描述标题"`
	Introduction string `json:"introduction" comment:"视频素材描述简介"`
}

func (p *AddMaterialRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
