package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type MassSendResponse struct {
	vo.Base
	Type      string `json:"type" comment:"媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb），次数为news，即图文消息"`
	MsgId     int    `json:"msg_id" comment:"消息发送任务的ID"`
	MsgDataId int    `json:"msg_data_id" comment:"消息的数据ID，，该字段只有在群发图文消息时，才会出现。可以用于在图文分析数据接口中，获取到对应的图文消息的数据，是图文分析数据接口中的msgid字段中的前半部分，详见图文分析数据接口中的msgid字段的介"`
}
