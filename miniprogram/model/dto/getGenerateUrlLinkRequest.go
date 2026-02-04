package dto

import (
	"encoding/json"
	"github.com/yijiacode188/wxSDK/miniprogram/consts"
)

type CloudBase struct {
	Env           string `json:"env" comment:"云开发环境"`
	Domain        string `json:"domain" comment:"静态网站自定义域名，不填则使用默认域名"`
	Path          string `json:"path" comment:"云开发静态网站 H5 页面路径，不可携带 query"`
	Query         string `json:"query" comment:"云开发静态网站 H5 页面 query 参数，最大 1024 个字符，只支持数字，大小写英文以及部分特殊字符"`
	ResourceAppId string `json:"resource_appid" comment:"第三方批量代云开发时必填，表示创建该 env 的 appid "`
}
type GetGenerateUrlLinkRequest struct {
	Path           string            `json:"path" comment:"通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面"`
	Query          string            `json:"query" comment:"通过 URL Link 进入小程序时的query，最大1024个字符，只支持数字，大小写英文以及部分特殊字"`
	ExpireType     int               `json:"expire_type" comment:"默认值0.小程序 URL Link 失效类型，失效时间：0，失效间隔天数：1"`
	ExpireTime     int               `json:"expire_time" comment:"	到期失效的 URL Link 的失效时间，为 Unix 时间戳。生成的到期失效 URL Link 在该时间前有效。最长有效期为30天。expire_type 为 0 必填"`
	ExpireInterval int               `json:"expire_interval" comment:"到期失效的URL Link的失效间隔天数。生成的到期失效URL Link在该间隔时间到达前有效。最长间隔天数为30天。expire_type 为 1 必填"`
	CloudBase      *CloudBase        `json:"cloud_base"`
	EnvVersion     consts.EnvVersion `json:"env_version" comment:"要打开的小程序版本。正式版为 release,体验版为 trial 开发版为 develop,默认为正式版"`
}

func (p *GetGenerateUrlLinkRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
