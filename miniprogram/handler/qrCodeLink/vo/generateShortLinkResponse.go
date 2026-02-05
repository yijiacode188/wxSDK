package vo

import (
	"github.com/yijiacode188/wxSDK/miniprogram/model/vo"
)

type GetGenerateShortLinkResponse struct {
	vo.Base
	Link string `json:"link"`
}
