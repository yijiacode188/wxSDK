package dto

import "encoding/json"

type DelWxaNewTemplateRequest struct {
	PriTmplId string `json:"priTmplId" comment:"要删除的模板id"`
}

func (p *DelWxaNewTemplateRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
