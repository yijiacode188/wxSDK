package dto

import "encoding/json"

type FreePublishBatchGetRequest struct {
	Offset    int64 `json:"offset" comment:"从全部素材的该偏移位置开始返回，0表示从第一个素材返回"`
	Count     int64 `json:"count" comment:"返回素材的数量，取值在1到20之间"`
	NoContent int   `json:"no_content" comment:"1 表示不返回content字段，0表示正常返回，默认为0"`
}

func (p *FreePublishBatchGetRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
