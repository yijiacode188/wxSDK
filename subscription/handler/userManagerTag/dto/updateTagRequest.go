package dto

import "encoding/json"

type UpdateTagInfo struct {
	Id   int64  `json:"id" comment:"标签id，由微信分配"`
	Name string `json:"name" comment:"修改的标签名，UTF8编码"`
}
type UpdateTagRequest struct {
	Tag UpdateTagInfo `json:"tag" comment:"标签信息"`
}

func (p *UpdateTagRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
