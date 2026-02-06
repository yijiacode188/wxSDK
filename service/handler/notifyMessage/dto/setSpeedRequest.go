package dto

import "encoding/json"

type SetSpeedRequest struct {
	Speed int `json:"speed" comment:"群发速度的级别（0 为 80w/分钟、1 为 60w/分钟、2 为 45w/分钟、3 为 30w/分钟、4 为 10w/分钟）"`
}

func (p *SetSpeedRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
