package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type CreateTagResponse struct {
	vo.Base
	Tag struct {
		Id   int64  `json:"id" comment:"ID"`
		Name string `json:"name" comment:"标签名"`
	} `json:"tag" comment:"标签信息"`
}
