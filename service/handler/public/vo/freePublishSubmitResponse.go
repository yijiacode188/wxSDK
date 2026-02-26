package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type FreePublishSubmitResponse struct {
	vo.Base
	PublishId string `json:"publish_id" comment:"发布任务的id"`
	MsgDataId string `json:"msg_data_id" comment:"消息的数据ID"`
}
