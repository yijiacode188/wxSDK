package dto

import "encoding/json"

type DeleteConditionalMenuRequest struct {
	MenuId string `json:"menuid" comment:"菜单ID"`
}

func (p *DeleteConditionalMenuRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
