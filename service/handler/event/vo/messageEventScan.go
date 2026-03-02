package vo

// MessageEventScan  用户已关注时的扫码事件推送
type MessageEventScan struct {
	ToUserName   string `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64  `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string `xml:"MsgType" comment:"消息类型，event"`
	Event        string `xml:"Event" comment:"SCAN"`
	EventKey     string `xml:"EventKey" comment:"事件KEY值，为二维码的场景值ID"`
	Ticket       string `xml:"Ticket" comment:"二维码的ticket，可用来换取二维码图片"`
}
