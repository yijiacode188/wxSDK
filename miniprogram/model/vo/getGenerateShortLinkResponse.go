package vo

import "wxSDK/miniprogram/model/vo/base"

type GetGenerateShortLinkResponse struct {
	base.Base
	Link string `json:"link"`
}
