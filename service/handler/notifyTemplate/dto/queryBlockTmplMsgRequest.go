package dto

import "encoding/json"

type QueryBlockTmplMsgRequest struct {
	TmplMsgId string `json:"tmpl_msg_id" comment:"被拦截的模板消息id"`
	LargestId int    `json:"largest_id" comment:"上一页查询结果最大的id，用于翻页，第一次传0"`
	Limit     int    `json:"limit" comment:"单页查询的大小，最大100"`
}

func (p *QueryBlockTmplMsgRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
