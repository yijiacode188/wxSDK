package dto

import "encoding/xml"

type SendTextMessageRequest struct {
	ToUserName   string   `xml:"ToUserName" comment:"接收方账号（收到的OpenID）"`
	FromUserName string   `xml:"FromUserName" comment:"开发者微信号"`
	CreateTime   int64    `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string   `xml:"MsgType" comment:"消息类型，文本为text"`
	Content      string   `xml:"Content" comment:"回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）"`
	XMLName      xml.Name `xml:"xml"`
}

func (s *SendTextMessageRequest) ToXML() ([]byte, error) {
	return xml.Marshal(s)
}
