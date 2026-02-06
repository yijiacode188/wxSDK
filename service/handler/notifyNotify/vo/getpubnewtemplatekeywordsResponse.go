package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type Keyword struct {
	Kid     int    `json:"kid" comment:"关键词 id，选用模板时需要"`
	Name    string `json:"name" comment:"关键词内容"`
	Example string `json:"example" comment:"关键词内容对应的示例"`
	Rule    string `json:"rule" comment:"参数类型"`
}
type GetPubNewTemplateKeywordsResponse struct {
	vo.Base
	Count int       `json:"count" comment:"模版标题列表总数"`
	Data  []Keyword `json:"data" comment:"关键词列表"`
}
