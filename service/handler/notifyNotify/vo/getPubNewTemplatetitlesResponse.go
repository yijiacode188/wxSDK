package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type TemplateTitleItem struct {
	Tid        int    `json:"tid" comment:"模版标题 id"`
	Title      string `json:"title" comment:"模版标题"`
	Type       int    `json:"type" comment:"模版类型，2 为一次性订阅，3 为长期订阅"`
	CategoryId int    `json:"categoryId" comment:"模版所属类目 id"`
}
type GetPubNewTemplateTitlesResponse struct {
	vo.Base
	Count int                 `json:"count" comment:"模版标题列表总数"`
	Data  []TemplateTitleItem `json:"data" comment:"模板标题列表"`
}
