package dto

import "encoding/json"

type GetPhoneNumberRequest struct {
	Code string `json:"code"`
}

func (p *GetPhoneNumberRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
