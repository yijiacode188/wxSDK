package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type UploadNewsMsgResponse struct {
	vo.Base
	Type      string `json:"type" comment:"媒体文件类型，图文消息（news）"`
	MediaId   string `json:"media_id" comment:"图文消息上传后获取的唯一标识"`
	CreatedAt string `json:"created_at" comment:"媒体文件上传时间"`
}
