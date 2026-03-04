package dto

import "encoding/json"

type DraftBatchGetRequest struct {
	Offset    int64 `json:"offset" comment:"从全部素材的该偏移位置开始返回，0表示从第一个素材返回"`
	Count     int64 `json:"count" comment:"返回素材的数量，取值在1到20之间"`
	NoContent int   `json:"no_content" comment:"1 表示不返回 content 字段，0 表示正常返回，默认为 0"`
}

func (p *DraftBatchGetRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
