package dto

import "encoding/json"

type SendAllFilter struct {
	IsToAll string `json:"is_to_all" comment:"用于设定是否向全部用户发送，值为true或false，选择true该消息群发给所有用户，选择false可根据tag_id发送给指定群组的用户"`
	TagId   string `json:"tag_id" comment:"群发到的标签的tag_id，参见用户管理中用户分组接口，若is_to_all值为true，可不填写tag_id"`
}
type SendAllMpVide struct {
	MediaId string `json:"media_id" comment:"用于群发的图文消息的media_id"`
}
type SendAllRequest struct {
	Filter      SendAllFilter `json:"filter" comment:"用于设定图文消息的接收者"`
	MsgType     string        `json:"msgtype" comment:"群发的消息类型，图文消息为mpnews，文本消息为text，语音为voice，音乐为music，图片为image，视频为mpvideo，卡券为wxcard"`
	MpNews      MpNews        `json:"mpnews" comment:"图文消息"`
	Text        Text          `json:"text" comment:"文本消息"`
	Voice       Voice         `json:"voice" comment:"语音消息"`
	Images      Images        `json:"images" comment:"图片消息"`
	MpVideo     MpVideo       `json:"mpvideo" comment:"视频消息"`
	WXCard      WXCard        `json:"wxcard" comment:"卡券消息"`
	ClientMsgId string        `json:"clientmsgid" comment:"开发者侧群发msgid，长度限制32字节，如不填，则后台默认以群发范围和群发内容的摘要值做为 clientmsgid"`
}

func (p *SendAllRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
