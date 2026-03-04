package vo

import "encoding/xml"

// MessageImage 图片消息
type MessageImage struct {
	ToUserName   string   `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string   `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64    `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string   `xml:"MsgType" comment:"消息类型，文本为text"`
	PicUrl       string   `xml:"PicUrl" comment:"图片链接（由系统生成）"`
	MediaId      int64    `xml:"MediaId" comment:"图片消息媒体id，可以调用获取临时素材接口拉取数据。"`
	MsgId        int64    `xml:"MsgId" comment:"消息id，64位整型"`
	MsgDataId    int64    `xml:"MsgDataId" comment:"消息的数据ID（消息如果来自文章时才有）"`
	Idx          int      `xml:"Idx" comment:"多图文时第几篇文章，从1开始（消息如果来自文章时才有）"`
	XMLName      xml.Name `xml:"xml"`
}
