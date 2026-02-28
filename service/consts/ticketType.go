package consts

type TicketType string

const (
	TicketTypeJSSDK TicketType = "jsapi"   //为 js-sdk 凭证
	TicketTypeCard  TicketType = "wx_card" //为微信卡券凭证
)
