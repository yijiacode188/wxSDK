package vo

import "github.com/yijiacode188/wxSDK/service/model/vo"

type FetchShortenResponse struct {
	vo.Base
	LongData      string `json:"long_data" comment:"长信息"`
	CreateTime    int64  `json:"create_time" comment:"创建的时间戳"`
	ExpireSeconds int64  `json:"expire_seconds" comment:"剩余的过期秒数"`
}
