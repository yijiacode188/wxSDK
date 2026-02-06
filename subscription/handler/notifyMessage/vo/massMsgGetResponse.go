package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type MassMsgGetResponse struct {
	vo.Base
	MsgId     string `json:"msg_id" comment:"群发消息后返回的消息id"`
	MsgStatus string `json:"msg_status" comment:"消息发送后的状态，SEND_SUCCESS表示发送成功，SENDING表示发送中，SEND_FAIL表示发送失败，DELETE表示已删除"`
}
