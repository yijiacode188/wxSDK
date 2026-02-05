package vo

import (
	"github.com/yijiacode188/wxSDK/subscription/model/vo"
)

type GetApiDomainIpResponse struct {
	vo.Base
	IPList []string `json:"ip_list" comment:"微信服务器IP地址列表"`
}
