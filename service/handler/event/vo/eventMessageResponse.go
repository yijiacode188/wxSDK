package vo

import (
	"encoding/xml"
	"fmt"
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

	PicUrl     string `xml:"PicUrl" comment:"图片链接（由系统生成）"`
	MediaId    int64  `xml:"MediaId" comment:"图片消息媒体id，可以调用获取临时素材接口拉取数据。"`
	Format     string `xml:"Format" comment:"语音格式，如amr，speex等"`
	MediaId16K string `xml:"MediaId16K" comment:"16K采样率语音消息媒体id，可以调用获取临时素材接口拉取数据，返回16K采样率amr/speex语音。"`

	ThumbMediaId int64   `xml:"ThumbMediaId" comment:"视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。"`
	Location_X   float32 `xml:"Location_X" comment:"地理位置纬度"`
	Location_Y   float32 `xml:"Location_Y" comment:"地理位置经度"`
	Scale        float32 `xml:"Scale" comment:"地图缩放大小"`
	Label        string  `xml:"Label" comment:"地理位置信息"`

	Title       string `xml:"Title" comment:"消息标题"`
	Description string `xml:"Description" comment:"消息描述"`
	Url         string `xml:"Url" comment:"消息链接"`
	Ticket      string `xml:"Ticket" comment:"二维码的ticket，可用来换取二维码图片"`

	Latitude  float32  `xml:"Latitude" comment:"地理位置纬度"`
	Longitude float32  `xml:"Longitude" comment:"地理位置经度"`
	Precision float32  `xml:"Precision" comment:"地理位置精度"`
	XMLName   xml.Name `xml:"xml"`
}

// ToTextMsg 获取文本消息
func (e *EventMessageResponse) ToTextMsg() *MessageText {
	textMessage := &MessageText{}
	err := utils.AssignByTag(e, textMessage, []string{}, "xml")
	if err != nil {
		fmt.Println("出现了错误", err)
	}
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

// ToVideoMsg 获取视频消息
func (e *EventMessageResponse) ToVideoMsg() *MessageVideo {
	videoMessage := &MessageVideo{}
	utils.AssignByTag(e, videoMessage, []string{}, "xml")
	return videoMessage
}

// ToShortVideoMsg 获取短视频消息
func (e *EventMessageResponse) ToShortVideoMsg() *MessageShortVideo {
	shortVideoMessage := &MessageShortVideo{}
	utils.AssignByTag(e, shortVideoMessage, []string{}, "xml")
	return shortVideoMessage
}

// ToLocationMsg 获取地理位置信息
func (e *EventMessageResponse) ToLocationMsg() *MessageLocation {
	locationMessage := &MessageLocation{}
	utils.AssignByTag(e, locationMessage, []string{}, "xml")
	return locationMessage
}

// ToLinkMsg 获取链接消息
func (e *EventMessageResponse) ToLinkMsg() *MessageLink {
	linkMessage := &MessageLink{}
	utils.AssignByTag(e, linkMessage, []string{}, "xml")
	return linkMessage
}

// ToSubscribeEventMsg 获取订阅、取消订阅消息
func (e *EventMessageResponse) ToSubscribeEventMsg() *MessageEventSubscribe {
	subscribeMessage := &MessageEventSubscribe{}
	utils.AssignByTag(e, subscribeMessage, []string{}, "xml")
	return subscribeMessage
}

// ToScanEventMsg 获取扫码，用户已关注时的消息
func (e *EventMessageResponse) ToScanEventMsg() *MessageEventScan {
	scanEventMessage := &MessageEventScan{}
	utils.AssignByTag(e, scanEventMessage, []string{}, "xml")
	return scanEventMessage
}

// ToLocalEventMsg 获取上报地理位置事件
func (e *EventMessageResponse) ToLocalEventMsg() *MessageEventLocation {
	locationEventMessage := &MessageEventLocation{}
	utils.AssignByTag(e, locationEventMessage, []string{}, "xml")
	return locationEventMessage
}

// ToClickEventMsg 点击菜单拉取消息时的事件推送
func (e *EventMessageResponse) ToClickEventMsg() *MessageEventClick {
	clickEventMessage := &MessageEventClick{}
	utils.AssignByTag(e, clickEventMessage, []string{}, "xml")
	return clickEventMessage
}

// ToViewEventMsg 点击菜单跳转链接时的事件推送
func (e *EventMessageResponse) ToViewEventMsg() *MessageEventView {
	viewEventMessage := &MessageEventView{}
	utils.AssignByTag(e, viewEventMessage, []string{}, "xml")
	return viewEventMessage
}
