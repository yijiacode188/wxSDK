package dto

import "encoding/json"

type Music struct {
	MediaId string `json:"media_id" comment:"用于群发的图文消息的media_id"`
}
type Image struct {
	MediaId string `json:"media_id" comment:"用于群发的图文消息的media_id"`
}
type PreMpVideo struct {
	MediaId string `json:"media_id" comment:"用于群发的图文消息的media_id"`
}
type PreWXCard struct {
	CardId  string      `json:"card_id" comment:"卡券ID"`
	CardExt *PreCardExt `json:"card_ext" comment:"卡券属性"`
}
type PreCardExt struct {
	Code      string `json:"code" comment:"卡券code"`
	OpenId    string `json:"openid" comment:"用户openid"`
	Timestamp string `json:"timestamp" comment:"时间戳"`
	Signature string `json:"signature" comment:"签名"`
}
type PreviewRequest struct {
	ToUser   string      `json:"touser" comment:"接收消息用户对应该公众号的openid（touser、towxname二选一）"`
	ToWxName string      `json:"towxname" comment:"接收消息用户微信号，实现对微信号的预览（touser、towxname二选一）"`
	MsgType  string      `json:"msgtype" comment:"群发的消息类型，图文消息为mpnews，文本消息为text，语音为voice，音乐为music，图片为image，视频为mpvideo，卡券为wxcard"`
	MpNews   *MpNews     `json:"mpnews" comment:"图文消息"`
	Text     *Text       `json:"text" comment:"文本消息"`
	Voice    *Voice      `json:"voice" comment:"语音消息"`
	Music    *Music      `json:"music" comment:"音乐消息"`
	Image    *Image      `json:"image" comment:"图片消息"`
	MpVideo  *PreMpVideo `json:"mpvideo" comment:"视频消息"`
	WXCard   *PreWXCard  `json:"wxcard" comment:"卡券消息"`
}

func (p *PreviewRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
