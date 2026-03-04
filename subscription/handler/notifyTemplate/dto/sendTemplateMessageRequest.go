package dto

import "encoding/json"

type MiniProgram struct {
	AppID    string `json:"appid" comment:"小程序appid"`
	PagePath string `json:"pagepath" comment:"小程序跳转路径"`
}
type DataItem struct {
	Value any `json:"value"`
}
type TemplateData map[string]DataItem
type SendTemplateMessageRequest struct {
	ToUser      string       `json:"touser" comment:"接收者（用户）的 openid"`
	TemplateId  string       `json:"template_id" comment:"所需下发的订阅模板id"`
	Url         string       `json:"url" comment:"模板跳转链接（海外账号没有跳转能力,url 和 miniprogram 同时不填，无跳转，url 和 miniprogram 同时填写，优先跳转小程序）"`
	MiniProgram MiniProgram  `json:"miniprogram" comment:"跳转小程序时填写（url 和 miniprogram 同时不填，无跳转，page 和 miniprogram 同时填写，优先跳转小程序）"`
	Data        TemplateData `json:"data" comment:"模板内容，需根据模板给定的格式给出（参考注意事项），格式形如 { \"key1\": { \"value\": any }, \"key2\": { \"value\": any } }"`
	ClientMsgId string       `json:"client_msg_id" comment:"防重入id。对于同一个openid + client_msg_id, 只发送一条消息,10分钟有效,超过10分钟不保证效果。若无防重入需求，可不填"`
}

func (p *SendTemplateMessageRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
