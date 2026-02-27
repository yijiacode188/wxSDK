package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetBlackListResponse struct {
	vo.Base
	Total int64 `json:"total" comment:"用户总数"`
	Count int64 `json:"count" comment:"本次返回的用户数"`
	Data  struct {
		OpenId []string `json:"openid" comment:"openid 列表"`
	} `json:"data" comment:"用户数据"`
	NextOpenId string `json:"next_openid" comment:"本次列表后一位openid"`
}
