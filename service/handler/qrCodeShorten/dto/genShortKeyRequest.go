package dto

import "encoding/json"

type GenShortKeyRequest struct {
	LongData      string `json:"long_data" comment:"需要转换的长信息，不超过4KB"`
	ExpireSeconds int64  `json:"expire_seconds" comment:"过期秒数，最大值为2592000（即30天），默认为2592000"`
}

func (p *GenShortKeyRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
