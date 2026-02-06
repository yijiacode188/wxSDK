package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type PubTemplate struct {
	PriTmplId            string        `json:"priTmplId" comment:"添加至帐号下的模板 id，发送小程序订阅消息时所需"`
	Title                string        `json:"title" comment:"模版标题"`
	Content              string        `json:"content" comment:"模版内容"`
	Example              string        `json:"example" comment:"模板内容示例"`
	Type                 int           `json:"type" comment:"模版类型，2 为一次性订阅，3 为长期订阅"`
	KeywordEnumValueList []KeywordEnum `json:"keywordEnumValueList" comment:"枚举参数值范围"`
}
type KeywordEnum struct {
	KeywordCode   string `json:"keywordCode" comment:"枚举参数的 key"`
	EnumValueList []any  `json:"enumValueList" comment:"枚举参数值范围列表"`
}
type GetWXaPubNewTemplateResponse struct {
	vo.Base
	Data []PubTemplate `json:"data" comment:"模板列表"`
}
