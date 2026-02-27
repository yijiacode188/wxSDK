package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type GetTagFansResponse struct {
	vo.Base
	Count int64 `json:"count" comment:"本次获取的粉丝数量"`
	Data  struct {
		OpenId []string `json:"openid" comment:"粉丝 openid 列表"`
	} `json:"data" comment:"标签下粉丝数据"`
	NextOpenId string `json:"next_openid" comment:"拉取列表最后一个用户的openid"`
}
