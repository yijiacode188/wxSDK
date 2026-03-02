package vo

import "encoding/xml"

// MessageLocation 地理位置消息
type MessageLocation struct {
	ToUserName   string   `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string   `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64    `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string   `xml:"MsgType" comment:"消息类型，地理位置为location"`
	Location_X   float32  `xml:"Location_X" comment:"地理位置纬度"`
	Location_Y   float32  `xml:"Location_Y" comment:"地理位置经度"`
	Scale        float32  `xml:"Scale" comment:"地图缩放大小"`
	Label        string   `xml:"Label" comment:"地理位置信息"`
	MsgId        int64    `xml:"MsgId" comment:"消息id，64位整型"`
	MsgDataId    int64    `xml:"MsgDataId" comment:"消息的数据ID（消息如果来自文章时才有）"`
	Idx          int      `xml:"Idx" comment:"多图文时第几篇文章，从1开始（消息如果来自文章时才有）"`
	XMLName      xml.Name `xml:"xml"`
}
