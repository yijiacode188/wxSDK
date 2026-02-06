package dto

import "encoding/json"

type SubscribeMsgItem struct {
	Value string `json:"value"`
}
type SubscribeMsg map[string]SubscribeMsgItem
type SendNewSubscribeMsgRequest struct {
	TemplateId       string       `json:"template_id" comment:"所需下发的订阅模板id"`
	Page             string       `json:"page" comment:"点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转"`
	ToUser           string       `json:"touser" comment:"接收者（用户）的 openid"`
	Data             SubscribeMsg `json:"data" comment:"模板内容，格式形如"`
	MiniProgramState string       `json:"miniprogram_state" comment:"跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版"`
	Lang             string       `json:"lang" comment:"进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN"`
}

func (p *SendNewSubscribeMsgRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
