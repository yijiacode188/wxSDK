package dto

import "encoding/xml"

// SendImageMessageRequest 回复图片消息
type SendImageMessageRequest struct {
	ToUserName   string   `xml:"ToUserName" comment:"接收方账号（收到的OpenID）"`
	FromUserName string   `xml:"FromUserName" comment:"开发者微信号"`
	CreateTime   int64    `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string   `xml:"MsgType" comment:"消息类型，图片为image"`
	MediaId      int64    `xml:"MediaId" comment:"通过素材管理中的接口上传多媒体文件，得到的id。"`
	XMLName      xml.Name `xml:"xml"`
}

func (s *SendImageMessageRequest) ToXML() ([]byte, error) {
	return xml.Marshal(s)
}
