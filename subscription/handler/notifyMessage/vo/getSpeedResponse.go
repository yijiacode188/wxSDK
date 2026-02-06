package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type GetSpeedResponse struct {
	vo.Base
	Speed     int `json:"speed" comment:"群发速度的级别 0 80w/分钟 1 60w/分钟 2 45w/分钟 3 30w/分钟 4 10w/分钟"`
	RealSpeed int `json:"realspeed" comment:"群发速度的真实值 单位：万/分钟"`
}
