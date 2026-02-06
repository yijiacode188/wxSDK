package dto

import "encoding/json"

type DeleteTemplateRequest struct {
	TemplateId string `json:"template_id" comment:"公众账号下模板消息ID, 包括类目模板ID"`
}

func (p *DeleteTemplateRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
