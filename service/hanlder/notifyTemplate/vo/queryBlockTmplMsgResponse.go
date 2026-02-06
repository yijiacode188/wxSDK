package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type MsgInfo struct {
	Id            string `json:"id" comment:"记录唯一ID，用于多次查询的largest_id"`
	TmplMsgId     string `json:"tmpl_msg_id" comment:"被拦截的模板消息id"`
	Title         string `json:"title" comment:"模板消息的标题"`
	Content       string `json:"content" comment:"模板消息的内容"`
	SendTimestamp int    `json:"send_timestamp" comment:"下发的时间戳"`
	OpenId        string `json:"openid" comment:"下发目标用户的openid"`
}
type QueryBlockTmplMsgResponse struct {
	vo.Base
	MsgInfo MsgInfo `json:"msginfo" comment:"消息信息"`
}
