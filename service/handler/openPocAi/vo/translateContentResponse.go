package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type TranslateContentResponse struct {
	vo.Base
	FromContent string `json:"from_content" comment:"原文内容"`
	ToContent   string `json:"to_content" comment:"译文内容"`
}
