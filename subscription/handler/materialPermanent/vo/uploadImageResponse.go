package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type UploadImageResponse struct {
	vo.Base
	Url string `json:"url" comment:"图片URL"`
}
