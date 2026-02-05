package dto

import "encoding/json"

type TryMatchMenuRequest struct {
	UserId string `json:"user_id" comment:"用户OpenID或微信号"`
}

func (p *TryMatchMenuRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
