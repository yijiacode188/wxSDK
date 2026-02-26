package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type UploadTempMediaResponse struct {
	vo.Base
	Type     string `json:"type" comment:"媒体文件类型"`
	MediaId  string `json:"media_id" comment:"媒体文件标识"`
	CreateAt int64  `json:"create_at" comment:"上传时间戳"`
}
