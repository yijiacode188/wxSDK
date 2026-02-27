package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetTagIdListResponse struct {
	vo.Base
	TagIdList []int64 `json:"tagid_list" comment:"用户的标签列表"`
}
