package dto

import "encoding/json"

type DraftAddRequest struct {
	Article []ArticlesInfo `json:"article" comment:"图文素材集合"`
}

func (p *DraftAddRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
