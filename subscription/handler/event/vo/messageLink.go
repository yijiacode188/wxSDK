package vo

import "encoding/xml"

// MessageLink 链接消息
type MessageLink struct {
	ToUserName   string   `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string   `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64    `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string   `xml:"MsgType" comment:"消息类型，链接为link"`
	Title        string   `xml:"Title" comment:"消息标题"`
	Description  string   `xml:"Description" comment:"消息描述"`
	Url          string   `xml:"Url" comment:"消息链接"`
	MsgId        int64    `xml:"MsgId" comment:"消息id，64位整型"`
	MsgDataId    int64    `xml:"MsgDataId" comment:"消息的数据ID（消息如果来自文章时才有）"`
	Idx          int      `xml:"Idx" comment:"多图文时第几篇文章，从1开始（消息如果来自文章时才有）"`
	XMLName      xml.Name `xml:"xml"`
}
