package dto

import "encoding/json"

type GetRidInfoRequest struct {
	Rid string `json:"rid" comment:"调用接口报错返回的rid"`
}

func (p *GetRidInfoRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
