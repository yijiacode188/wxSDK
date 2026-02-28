package vo

import (
	"encoding/xml"
	"github.com/yijiacode188/wxSDK/utils"
)

type EventMessageResponse struct {
	ToUserName   string `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	Content      string `xml:"Content"`
	CreateTime   int64  `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string `xml:"MsgType" comment:"消息类型，文本为text"`
	MsgId        int64  `xml:"MsgId" comment:"消息id，64位整型"`
	MsgDataId    int64  `xml:"MsgDataId" comment:"消息的数据ID（消息如果来自文章时才有）"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`

	Encrypt string `json:"Encrypt"`
	Idx     int    `xml:"Idx" comment:"多图文时第几篇文章，从1开始（消息如果来自文章时才有）"`

	PicUrl  string `xml:"PicUrl"`
	MediaId int64  `xml:"MediaId"`

	XMLName xml.Name `xml:"xml"`
}

// ToTextMsg 获取文本消息
func (e *EventMessageResponse) ToTextMsg() *MessageText {
	textMessage := &MessageText{}
	utils.AssignByTag(e, textMessage, []string{}, "xml")
	return textMessage
}

// ToImageMsg 获取图片消息
func (e *EventMessageResponse) ToImageMsg() *MessageImage {
	imageMessage := &MessageImage{}
	utils.AssignByTag(e, imageMessage, []string{}, "xml")
	return imageMessage
}

// ToVoiceMsg 获取语音消息
func (e *EventMessageResponse) ToVoiceMsg() *MessageVoice {
	voiceMessage := &MessageVoice{}
	utils.AssignByTag(e, voiceMessage, []string{}, "xml")
	return voiceMessage
}
