package vo

import (
	"wxSDK/miniprogram/consts"
	"wxSDK/miniprogram/model/vo/base"
)

type UrlLinkInfo struct {
	AppID      string            `json:"appid" comment:"小程序 appid"`
	Path       string            `json:"path" comment:"小程序页面路径"`
	Query      string            `json:"query" comment:"小程序页面query"`
	CreateTime int               `json:"create_time" comment:"创建时间，为 Unix 时间戳"`
	ExpireTime int               `json:"expire_time" comment:"到期失效时间，为 Unix 时间戳，0 表示永久生效"`
	EnvVersion consts.EnvVersion `json:"env_version" comment:"要打开的小程序版本"`
}
type QuotaInfo struct {
	RemainVisitQuota int `json:"remain_visit_quota" comment:"URL Scheme（加密+明文）/加密 URL Link 单天剩余访问次数"`
}
type GetQueryUrlLinkResponse struct {
	base.Base
	UrlLinkInfo UrlLinkInfo `json:"url_link_info" comment:"url_link 配置"`
	QuotaInfo   QuotaInfo   `json:"quota_info" comment:"quota 配置"`
}
