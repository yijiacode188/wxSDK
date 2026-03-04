package vo

import "github.com/yijiacode188/wxSDK/subscription/model/vo"

type DraftCountResponse struct {
	vo.Base
	TotalCount int64 `json:"total_count" comment:"草稿的总数"`
}
