package vo

// MessageEventLocation 上报地理位置事件
type MessageEventLocation struct {
	ToUserName   string  `xml:"ToUserName" comment:"开发者微信号"`
	FromUserName string  `xml:"FromUserName" comment:"发送方账号（一个OpenID）"`
	CreateTime   int64   `xml:"CreateTime" comment:"消息创建时间 （整型）"`
	MsgType      string  `xml:"MsgType" comment:"消息类型，event"`
	Event        string  `xml:"Event" comment:"事件类型，LOCATION"`
	Latitude     float32 `xml:"Latitude" comment:"地理位置纬度"`
	Longitude    float32 `xml:"Longitude" comment:"地理位置经度"`
	Precision    float32 `xml:"Precision" comment:"地理位置精度"`
}
