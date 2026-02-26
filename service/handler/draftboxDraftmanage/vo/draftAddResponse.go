package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type DraftAddResponse struct {
	vo.Base
	MediaId string `json:"media_id" comment:"上传后的获取标志(不超过128字符)"`
}
