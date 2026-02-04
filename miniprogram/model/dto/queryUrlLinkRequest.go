package dto

import "encoding/json"

type QueryUrlLinkRequest struct {
	UrlLink   string `json:"url_link" comment:"小程序加密 url_link"`
	QueryType int    `json:"query_type" comment:"查询类型。默认值0，查询 url_link 信息：0， 查询每天剩余访问次数：1"`
}

func (p *QueryUrlLinkRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
