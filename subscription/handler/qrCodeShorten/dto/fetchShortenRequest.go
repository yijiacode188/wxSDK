package dto

import "encoding/json"

type FetchShortenRequest struct {
	ShortKey string `json:"short_key" comment:"短key"`
}

func (p *FetchShortenRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
