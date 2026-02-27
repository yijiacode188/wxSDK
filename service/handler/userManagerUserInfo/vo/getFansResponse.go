package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetFansResponse struct {
	vo.Base
	Total int64 `json:"total" comment:"关注该公众账号的总用户数"`
	Count int64 `json:"count" comment:"拉取的OPENID个数，最大值为10000"`
	Data  struct {
		OpenId []string `json:"openid" comment:"用户唯一ID"`
	} `json:"data" comment:"列表数据，OPENID的列表"`
	NextOpenId string `json:"next_openid" comment:"拉取列表后一个用户的OPENID"`
}
