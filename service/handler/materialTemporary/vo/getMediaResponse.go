package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetMediaResponse struct {
	vo.Base
	VideoUrl string `json:"video_url" comment:"视频消息素材下载地址"`
}
