package vo

import (
	"github.com/yijiacode188/wxSDK/miniprogram/model/vo"
)

type GenerateUrlLinkResponse struct {
	vo.Base
	UrlLink string `json:"url_link"`
}
