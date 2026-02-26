package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type AddMaterialResponse struct {
	vo.Base
	MediaId string `json:"media_id" comment:"新增的永久素材media_id"`
	Url     string `json:"url" comment:"图片素材URL(仅图片返回)"`
}
