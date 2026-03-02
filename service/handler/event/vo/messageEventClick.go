package vo

// MessageEventClick 自定义菜单事件
type MessageEventClick struct {
	ToUserName   string `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64  `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string `xml:"MsgType" comment:"消息类型，event"`
	Event        string `xml:"Event" comment:"事件类型，CLICK"`
	EventKey     string `xml:"EventKey" comment:"事件KEY值，与自定义菜单接口中KEY值对应"`
}
