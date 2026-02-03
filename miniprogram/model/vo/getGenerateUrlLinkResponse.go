package vo

import "wxSDK/miniprogram/model/vo/base"

type GetGenerateUrlLinkResponse struct {
	base.Base
	UrlLink string `json:"url_link"`
}
