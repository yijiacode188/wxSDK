package vo

import "encoding/xml"

// MessageVoice 语音消息
type MessageVoice struct {
	ToUserName   string   `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string   `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64    `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string   `xml:"MsgType" comment:"消息类型，文本为text"`
	MediaId      int64    `xml:"MediaId" comment:"语音消息媒体id，可以调用获取临时素材接口拉取数据，Format为amr时返回8K采样率amr语音。"`
	Format       string   `xml:"Format" comment:"语音格式，如amr，speex等"`
	MsgId        int64    `xml:"MsgId" comment:"消息id，64位整型"`
	MsgDataId    int64    `xml:"MsgDataId" comment:"消息的数据ID（消息如果来自文章时才有）"`
	Idx          int      `xml:"Idx" comment:"多图文时第几篇文章，从1开始（消息如果来自文章时才有）"`
	MediaId16K   int64    `xml:"MediaId16K" comment:"16K采样率语音消息媒体id，可以调用获取临时素材接口拉取数据，返回16K采样率amr/speex语音。"`
	XMLName      xml.Name `xml:"xml"`
}
