package dto

import "encoding/json"

type MpNews struct {
	MediaId string `json:"media_id" comment:"用于群发的图文消息的media_id"`
}
type Text struct {
	Content string `json:"content" comment:"文本内容"`
}
type Voice struct {
	MediaId string `json:"media_id" comment:"用于群发的图文消息的media_id"`
}
type Images struct {
	MediaIds           []string `json:"media_ids" comment:"用于群发的图文消息的media_id列表"`
	ReCommend          string   `json:"recommend" comment:"推荐语"`
	Title              string   `json:"title" comment:"标题"`
	NeedOpenComment    int      `json:"need_open_comment" comment:"开启评论（1-开启，0-关闭）"`
	OnlyFansCanComment int      `json:"only_fans_can_comment" comment:"只有粉丝能评论（1-开启，0-关闭）"`
}
type MpVideo struct {
	MediaId     string `json:"media_id" comment:"用于群发的图文消息的media_id"`
	Title       string `json:"title" comment:"标题"`
	Description string `json:"description" comment:"描述"`
}
type WXCard struct {
	CardId string `json:"card_id" comment:"卡券 ID"`
}
type MassSendRequest struct {
	ToUser            []string `json:"touser" comment:"填写图文消息的接收者，一串OpenID列表，OpenID最少2个，最多10000个"`
	MsgType           string   `json:"msgtype" comment:"群发的消息类型，图文消息为mpnews，文本消息为text，语音为voice，音乐为music，图片为image，视频为video，卡券为wxcard"`
	SendIgnoreReprint int      `json:"send_ignore_reprint" comment:"图文消息被判定为转载时，是否继续群发。 1为继续群发（转载），0为停止群发。 该参数默认为0。"`
	MpNews            *MpNews  `json:"mpnews" comment:"图文消息"`
	Text              *Text    `json:"text" comment:"文本消息"`
	Voice             *Voice   `json:"voice" comment:"语音消息"`
	Images            *Images  `json:"images" comment:"图片消息"`
	MpVideo           *MpVideo `json:"mpvideo" comment:"视频消息"`
	WXCard            *WXCard  `json:"wxcard" comment:"卡券消息"`
	ClientMsgId       string   `json:"clientmsgid" comment:"开发者侧群发msgid，长度限制32字节，如不填，则后台默认以群发范围和群发内容的摘要值做为 clientmsgid"`
}

func (p *MassSendRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
