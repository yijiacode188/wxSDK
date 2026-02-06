package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type SendTemplateMessageResponse struct {
	vo.Base
	MsgId int `json:"msgid" comment:"消息id"`
}
