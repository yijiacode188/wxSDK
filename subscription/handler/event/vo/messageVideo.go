package vo

import "encoding/xml"

// MessageVideo 视频消息
type MessageVideo struct {
	ToUserName   string   `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string   `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64    `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string   `xml:"MsgType" comment:"视频为video"`
	MediaId      int64    `xml:"MediaId" comment:"视频消息媒体id，可以调用获取临时素材接口拉取数据。"`
	ThumbMediaId int64    `xml:"ThumbMediaId" comment:"视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。"`
	MsgId        int64    `xml:"MsgId" comment:"消息id，64位整型"`
	MsgDataId    int64    `xml:"MsgDataId" comment:"消息的数据ID（消息如果来自文章时才有）"`
	Idx          int      `xml:"Idx" comment:"多图文时第几篇文章，从1开始（消息如果来自文章时才有）"`
	XMLName      xml.Name `xml:"xml"`
}
