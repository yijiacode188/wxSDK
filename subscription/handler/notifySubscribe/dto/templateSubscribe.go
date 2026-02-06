package dto

import "encoding/json"

type MiniProgram struct {
	AppID    string `json:"appid" comment:"小程序appid"`
	PagePath string `json:"pagepath" comment:"小程序页面路径"`
}
type SubscribeContent struct {
	Value string `json:"value" comment:"消息文本(200字内)"`
	Color string `json:"color" comment:"字体颜色"`
}
type SubscribeData struct {
	Content SubscribeContent `json:"content" comment:"内容信息"`
}
type TemplateSubscribeRequest struct {
	ToUser      string         `json:"touser" comment:"接收消息的用户openid"`
	TemplateId  string         `json:"template_id" comment:"订阅消息模板ID"`
	Url         string         `json:"url" comment:"点击消息跳转链接(需ICP备案)"`
	MiniProgram *MiniProgram   `json:"miniprogram" comment:"跳小程序配置"`
	Scene       string         `json:"scene" comment:"订阅场景值"`
	Title       string         `json:"title" comment:"消息标题(15字以内)"`
	Data        *SubscribeData `json:"data" comment:"内容"`
}

func (p *TemplateSubscribeRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
