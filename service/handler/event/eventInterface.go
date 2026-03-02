package event

import "github.com/yijiacode188/wxSDK/service/handler/event/vo"

type MessageEvent interface {
	GetTimestamp() string
	GetNonce() string
	GetSignature() string
	GetMsgSignature() string
	//CallBackTextMessage 处理普通文本消息
	CallBackTextMessage(text *vo.MessageText)
	//CallBackImageMessage 处理图片消息
	CallBackImageMessage(image *vo.MessageImage)
	//CallBackVoiceMessage 处理语音消息
	CallBackVoiceMessage(voice *vo.MessageVoice)
	//CallBackVideoMessage 处理视频消息
	CallBackVideoMessage(video *vo.MessageVideo)
	//CallBackShortVideoMessage 处理小视频消息
	CallBackShortVideoMessage(shortVideo *vo.MessageShortVideo)
	//CallBackLocationMessage 处理地理位置消息
	CallBackLocationMessage(location *vo.MessageLocation)
	//CallBackLinkMessage 处理链接消息
	CallBackLinkMessage(link *vo.MessageLink)

	//CallBackSubscribeEvent 处理关注/取消关注事件
	CallBackSubscribeEvent(subscribeEvent *vo.MessageEventSubscribe)
	//CallBackScanEvent 处理用户已关注时的扫码事件推送
	CallBackScanEvent(scanEvent *vo.MessageEventScan)
	//CallBackLocationEvent 处理上报地理位置事件
	CallBackLocationEvent(locationEvent *vo.MessageEventLocation)
	//CallBackClickEvent 处理自定义菜单事件
	CallBackClickEvent(clickEvent *vo.MessageEventClick)
	//CallBackViewEvent 处理点击菜单跳转链接时的事件推送
	CallBackViewEvent(viewEvent *vo.MessageEventView)
}
