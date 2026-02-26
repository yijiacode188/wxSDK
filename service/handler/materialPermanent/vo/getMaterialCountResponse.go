package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetMaterialCountResponse struct {
	vo.Base
	VoiceCount int64 `json:"voice_count" comment:"语音总数量"`
	VideoCount int64 `json:"video_count" comment:"视频总数量"`
	ImageCount int64 `json:"image_count" comment:"图片总数量"`
	NewsCount  int64 `json:"news_count" comment:"图文总数量"`
}
