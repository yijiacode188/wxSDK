package dto

import "encoding/json"

type DeleteTagInfo struct {
	Id int64 `json:"id" comment:"标签ID"`
}
type DeleteTagRequest struct {
	Tag DeleteTagInfo `json:"tag" comment:"标签信息"`
}

func (p *DeleteTagRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
