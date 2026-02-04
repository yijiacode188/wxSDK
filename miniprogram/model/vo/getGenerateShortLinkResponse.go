package vo

import "github.com/yijiacode188/wxSDK/miniprogram/model/vo/base"

type GetGenerateShortLinkResponse struct {
	base.Base
	Link string `json:"link"`
}
