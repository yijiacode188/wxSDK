package vo

type MessageEventSubscribe struct {
	ToUserName   string `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64  `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string `xml:"MsgType" comment:"消息类型，event"`
	Event        string `xml:"Event" comment:"事件类型，subscribe(订阅)、unsubscribe(取消订阅)"`
	EventKey     string `xml:"EventKey" comment:"事件KEY值，qrscene_为前缀，后面为二维码的场景值ID"`
	Ticket       string `xml:"Ticket" comment:"二维码的ticket，可用来换取二维码图片"`
}
