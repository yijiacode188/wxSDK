package dto

import "encoding/json"

type FreePublishGetRequest struct {
	PublishId string `json:"publish_id" comment:"发布任务id"`
}

func (p *FreePublishGetRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
