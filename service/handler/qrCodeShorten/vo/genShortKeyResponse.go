package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GenShortKeyResponse struct {
	vo.Base
	ShortKey string `json:"short_key" comment:"短key，15字节，base62编码(0-9/a-z/A-Z)"`
}
