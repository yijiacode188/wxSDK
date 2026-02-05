package dto

import (
	"encoding/json"
	"github.com/yijiacode188/wxSDK/subscription/consts"
)

type Button struct {
	Type      consts.ButtonType `json:"type" comment:"菜单的响应动作类型（与 sub_button 互斥）"`
	Name      string            `json:"name" comment:"菜单标题，不超过16个字节，子菜单不超过60个字节"`
	Key       string            `json:"key" comment:"菜单KEY值，用于消息接口推送，不超过128字节。click等点击类型必须。"`
	Url       string            `json:"url" comment:"网页链接，用户点击菜单可打开链接，不超过1024字节。 type为miniprogram时，不支持小程序的老版本客户端将打开本url。view、miniprogram类型必填。"`
	MediaId   string            `json:"media_id" comment:"调用新增永久素材接口返回的合法media_id。media_id类型和view_limited类型必须"`
	AppID     string            `json:"appid" comment:"小程序的appid（仅认证的账号可配置），miniprogram类型必填；且注意 appid 是小写的"`
	PagePath  string            `json:"pagepath" comment:"小程序的页面路径，miniprogram类型必须"`
	ArticleId string            `json:"article_id" comment:"发布后获得的合法 article_id，article_id类型和article_view_limited类型必须"`
	SubButton []SubButton       `json:"sub_button" comment:"二级菜单结构体数组"`
}
type SubButton struct {
	Type      consts.ButtonType `json:"type" comment:"菜单的响应动作类型（与 sub_button 互斥）"`
	Name      string            `json:"name" comment:"菜单标题，不超过16个字节，子菜单不超过60个字节"`
	Key       string            `json:"key" comment:"菜单KEY值，用于消息接口推送，不超过128字节。click等点击类型必须。"`
	Url       string            `json:"url" comment:"网页链接，用户点击菜单可打开链接，不超过1024字节。 type为miniprogram时，不支持小程序的老版本客户端将打开本url。view、miniprogram类型必填。"`
	MediaId   string            `json:"media_id" comment:"调用新增永久素材接口返回的合法media_id。media_id类型和view_limited类型必须"`
	AppID     string            `json:"appid" comment:"小程序的appid（仅认证的账号可配置），miniprogram类型必填；且注意 appid 是小写的"`
	PagePath  string            `json:"pagepath" comment:"小程序的页面路径，miniprogram类型必须"`
	ArticleId string            `json:"article_id" comment:"发布后获得的合法 article_id，article_id类型和article_view_limited类型必须"`
}
type CreateCustomMenuRequest struct {
	Button []Button `json:"button" comment:"一级菜单结构体数组"`
}

func (p *CreateCustomMenuRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
