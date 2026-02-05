package vo

import (
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
)

type GetApiQuotaResponse struct {
	vo.Base
	Quota struct {
		DailyLimit int `json:"daily_limit" comment:"当天该账号可调用该接口的次数"`
		Used       int `json:"used" comment:"当天已经调用的次数"`
		Remain     int `json:"remain" comment:"当天剩余调用次数"`
	} `json:"quota" comment:"quota详情"`
	RateLimit struct {
		CallCount     int `json:"call_count" comment:"周期内可调用数量，单位 次"`
		RefreshSecond int `json:"refresh_second" comment:"更新周期，单位 秒"`
	} `json:"rate_limit" comment:"普通调用频率限制"`
	ComponentRateLimit struct {
		CallCount     int `json:"call_count" comment:"周期内可调用数量，单位 次"`
		RefreshSecond int `json:"refresh_second" comment:"更新周期，单位 秒"`
	} `json:"component_rate_limit" comment:"代调用频率限制"`
}
