package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetTagsResponse struct {
	vo.Base
	Tags []struct {
		Id    int64  `json:"id" comment:"标签id，由微信分配"`
		Name  string `json:"name" comment:"标签名，UTF8编码"`
		Count int64  `json:"count" comment:"标签下粉丝数"`
	} `json:"tags" comment:"标签列表"`
}
