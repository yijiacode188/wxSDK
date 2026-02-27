package dto

import "encoding/json"

type TagInfo struct {
	Name string `json:"name" comment:"标签名（30个字符以内）"`
}
type CreateTagRequest struct {
	Tag TagInfo `json:"tag" comment:"标签信息"`
}

func (p *CreateTagRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
